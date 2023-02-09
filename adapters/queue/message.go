package queue

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"

	"gitlab.id.vin/gami/ps2-gami-common/utils"

	uuid "github.com/satori/go.uuid"
)

// KafkaMessage - message info
type KafkaMessage struct {
	Topic string
	Key   []byte
	Value []byte
}

// ProducerConfig config
type ProducerConfig struct {
	SeedBrokers    []string
	FlushConfig    *FlushConfig
	ClientCertFile string
	ClientKeyFile  string
	ClientCAFile   string
}
type FlushConfig struct {
	NumFlushMessages int
	MaxMessage       int
	FlushFrequency   time.Duration
}

// ConsumerConfig config
type ConsumerConfig struct {
	SeedBrokers     []string
	ConsumerGroupID string
	Topic           []string
	InitialOffset   InitOffsetType
	ClientCertFile  string
	ClientKeyFile   string
	ClientCAFile    string
}

// ConsumerMessage message struct
type ConsumerMessage struct {
	Key       []byte
	Value     []byte
	Topic     string
	Partition int32
	Offset    int64
	Timestamp time.Time
}

func NewTLSConfig(clientCertFile, clientKeyFile, caCertFile string) (*tls.Config, error) {
	tlsConfig := tls.Config{}

	// Load client cert
	cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return &tlsConfig, err
	}
	tlsConfig.Certificates = []tls.Certificate{cert}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		return &tlsConfig, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig.RootCAs = caCertPool

	tlsConfig.BuildNameToCertificate()
	return &tlsConfig, err
}

// KafkaMetaData info
type KafkaMetaData struct {
	ID         string      `json:"id"`
	RefID      string      `json:"ref_id"`
	Event      string      `json:"event"`
	Timestamps int64       `json:"timestamp"`
	Payload    interface{} `json:"payloads"`
}

func GetKafkaMessage(payload interface{}, refID, event string) KafkaMetaData {
	return KafkaMetaData{
		ID:         uuid.NewV4().String(),
		RefID:      refID,
		Event:      event,
		Timestamps: utils.ConvertTimeToMiliSeconds(time.Now().UTC()),
		Payload:    payload,
	}
}

type KafkaAttributeMetaData struct {
	ID         string           `json:"id"`
	Event      string           `json:"event"`
	UserID     string           `json:"user_id"`
	EventTime  int64            `json:"event_time"`
	RequestID  string           `json:"request_id"`
	PayloadID  string           `json:"payload_id"`
	ProducerID string           `json:"producer_id"`
	Payload    AttributePayload `json:"payload"`
}

type AttributePayload struct {
	UserID    string      `json:"user_id"`
	Attribute string      `json:"attribute"`
	Tags      interface{} `json:"tags"`
}

func GetKafkaAttributeMessage(tags interface{}, userID, reqID, event, collection, payloadID, producerID string) KafkaAttributeMetaData {
	return KafkaAttributeMetaData{
		ID:         uuid.NewV4().String(),
		Event:      event,
		UserID:     userID,
		EventTime:  utils.ConvertTimeToMiliSeconds(time.Now().UTC()),
		RequestID:  reqID,
		PayloadID:  payloadID,
		ProducerID: producerID,
		Payload: AttributePayload{
			Tags:      tags,
			UserID:    userID,
			Attribute: collection,
		},
	}
}
