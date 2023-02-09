package queue

import (
	"time"

	"gitlab.id.vin/gami/ps2-gami-common/logger"

	"github.com/Shopify/sarama"
)

// ProducerType type
type ProducerType string

// KafkaProducer is interface of kafka producer
type KafkaProducer interface {
	SendMessage(*KafkaMessage)
	SendMessageAsync(message *KafkaMessage)
	SendMessageSync(message *KafkaMessage) error
	Close()
}

type kafkaProducer struct {
	asyncProducer sarama.AsyncProducer
	syncProducer  sarama.SyncProducer
	done          chan bool
}

// NewKafkaProducer init Producer
func NewKafkaProducer(cf ProducerConfig, withTLS ...bool) (KafkaProducer, error) {
	if cf.FlushConfig == nil {
		cf.FlushConfig = &FlushConfig{
			NumFlushMessages: 100,
			MaxMessage:       1000,
			FlushFrequency:   10 * time.Millisecond,
		}
	}
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.Flush.Messages = cf.FlushConfig.NumFlushMessages
	config.Producer.Flush.Frequency = cf.FlushConfig.FlushFrequency
	config.Producer.Flush.MaxMessages = cf.FlushConfig.MaxMessage
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V1_0_0_0

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

	asyncProducer, err := sarama.NewAsyncProducer(cf.SeedBrokers, config)
	if err != nil {
		return nil, err
	}

	syncProducer, err := sarama.NewSyncProducer(cf.SeedBrokers, config)
	if err != nil {
		return nil, err
	}

	kafkaProducer := &kafkaProducer{
		asyncProducer: asyncProducer,
		syncProducer:  syncProducer,
		done:          make(chan bool),
	}

	go func() {
		for {
			select {
			case <-asyncProducer.Successes():
			case err := <-asyncProducer.Errors():
				logger.Errorf("send kafka message error %v", err)
			case <-kafkaProducer.done:
				return
			}
		}
	}()

	return kafkaProducer, nil
}

// SendMessageSync Async
func (p *kafkaProducer) SendMessageAsync(m *KafkaMessage) {
	msg := &sarama.ProducerMessage{
		Topic:     m.Topic,
		Key:       sarama.ByteEncoder(m.Key),
		Value:     sarama.ByteEncoder(m.Value),
		Timestamp: time.Now(),
	}

	p.asyncProducer.Input() <- msg
}

// SendMessageSync Info
func (p *kafkaProducer) SendMessageSync(m *KafkaMessage) error {
	msg := &sarama.ProducerMessage{
		Topic:     m.Topic,
		Key:       sarama.ByteEncoder(m.Key),
		Value:     sarama.ByteEncoder(m.Value),
		Timestamp: time.Now(),
	}

	_, _, err := p.syncProducer.SendMessage(msg)
	return err
}

// SendMessage send message to topic
func (p *kafkaProducer) SendMessage(m *KafkaMessage) {
	msg := &sarama.ProducerMessage{
		Topic:     m.Topic,
		Key:       sarama.ByteEncoder(m.Key),
		Value:     sarama.ByteEncoder(m.Value),
		Timestamp: time.Now(),
	}

	p.asyncProducer.Input() <- msg
}

// Close topic
func (p *kafkaProducer) Close() {
	p.asyncProducer.AsyncClose()
	err := p.syncProducer.Close()
	if err != nil {
		logger.Errorf("Close producer")
	}

	p.done <- true

}
