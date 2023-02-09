package queue

import (
	"context"
	"time"

	"gitlab.id.vin/gami/gami-common/logger"

	"github.com/Shopify/sarama"
)

// InitOffsetType define const
type InitOffsetType int64

// InitOffsetNewest const
const (
	InitOffsetNewest InitOffsetType = -1
	InitOffsetOldest InitOffsetType = -2
)

// KafkaConsumerHandlerFunc handle ...
type KafkaConsumerHandlerFunc func(message *ConsumerMessage)

type consumerGroupHandler struct {
	handleFunc KafkaConsumerHandlerFunc
}

// Setup ...
func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

// Cleanup - clean
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim - claim message
func (cg consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		cg.handleFunc(&ConsumerMessage{
			Key:       msg.Key,
			Value:     msg.Value,
			Topic:     msg.Topic,
			Partition: msg.Partition,
			Offset:    msg.Offset,
			Timestamp: msg.Timestamp,
		})
		sess.MarkMessage(msg, "")
	}
	return nil
}

// ConsumerHandler interface handler
type ConsumerHandler interface {
	HandlerFunc(*ConsumerMessage)
	Close()
}

// KafkaConsumer init consumer
type KafkaConsumer struct {
	client        sarama.Client
	consumerGroup sarama.ConsumerGroup
	Topic         []string
	Handler       ConsumerHandler
	running       bool
}

// NewKafkaConsumer - init consumer
func NewKafkaConsumer(cf ConsumerConfig, withTLS ...bool) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	config.Consumer.Return.Errors = true
	// edit configuration for issue consumer error when increase or remove pod with same consumer id
	config.Consumer.MaxProcessingTime = 500 * time.Millisecond
	config.Consumer.Group.Session.Timeout = 20 * time.Second
	config.Consumer.Group.Heartbeat.Interval = 6 * time.Second

	switch cf.InitialOffset {
	case InitOffsetNewest:
		config.Consumer.Offsets.Initial = sarama.OffsetNewest
	case InitOffsetOldest:
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	if len(withTLS) == 0 || (len(withTLS) > 0 && withTLS[0]) {
		tlsConfig, err := NewTLSConfig(
			cf.ClientCertFile,
			cf.ClientKeyFile,
			cf.ClientCAFile)

		if err != nil {
			logger.Fatalf(err, "Cannot load ssl file")
		}

		tlsConfig.InsecureSkipVerify = true
		config.Net.TLS.Enable = true
		config.Net.TLS.Config = tlsConfig
	}

	kkClient, err := sarama.NewClient(cf.SeedBrokers, config)
	if err != nil {
		return nil, err
	}

	kkConsumerGroup, err := sarama.NewConsumerGroupFromClient(cf.ConsumerGroupID, kkClient)
	if err != nil {
		return nil, err
	}

	consumer := &KafkaConsumer{
		client:        kkClient,
		consumerGroup: kkConsumerGroup,
		Topic:         cf.Topic,
	}

	return consumer, nil
}

// SetHandler controller consumer logic handler
func (c *KafkaConsumer) SetHandler(fn ConsumerHandler) {
	c.Handler = fn
}

// Start consumer
func (c *KafkaConsumer) Start() {
	// Track errors
	go func() {
		for err := range c.consumerGroup.Errors() {
			logger.Errorf("Consumer error: %v", err)
		}
	}()

	// Iterate over consumer sessions.
	c.running = true

	ctx := context.Background()
	for c.running {
		topics := c.Topic
		handler := consumerGroupHandler{handleFunc: c.Handler.HandlerFunc}

		err := c.consumerGroup.Consume(ctx, topics, handler)
		if err != nil {
			if !c.running {
				logger.Errorf("Consumer closed")
			} else {
				logger.Errorf("Consumer closed error: %v", err)
			}
		}
	}
}

// Close consumer
func (c *KafkaConsumer) Close() {
	c.running = false
	_ = c.consumerGroup.Close()
	c.Handler.Close()
}
