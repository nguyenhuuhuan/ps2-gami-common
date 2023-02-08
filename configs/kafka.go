package configs

// Kafka struct define
type Kafka struct {
	SeedBrokers    []string `default:"vep-kafka-1.int.vinid.dev:9093,vep-kafka-2.int.vinid.dev:9093,vep-kafka-3.int.vinid.dev:9093" envconfig:"KAFKA_SEEDBROKERS"`
	Topics         []string `default:"c1.promotion.gami-mission-hit.local" envconfig:"KAFKA_TOPIC"`
	GroupID        string   `default:"lottery" envconfig:"KAFKA_GROUPID"`
	InitialOffset  int64    `default:"-1" envconfig:"KAFKA_INITIAL_OFFSET"`
	ClientCertFile string   `default:"./keys/kafka/gamification-clients.int.vinid.dev-cert.pem" envconfig:"KAFKA_CLIENT_CERT"`
	ClientKeyFile  string   `default:"./keys/kafka/gamification-clients.int.vinid.dev-key.pem" envconfig:"KAFKA_CLIENT_KEY"`
	ClientCAFile   string   `default:"./keys/kafka/gamification-kafka-int-ca.pem" envconfig:"KAFKA_CLIENT_CA"`
	SSL            bool     `default:"true" envconfig:"KAFKA_SSL"`
}
