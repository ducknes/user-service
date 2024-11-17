package settings

type Config struct {
	Port      int      `json:"port"`
	Databases Database `json:"databases"`
}

type Database struct {
	MongoDB MongoSettings `json:"mongo"`
	Kafka   KafkaSettings `json:"kafka"`
}

type MongoSettings struct {
	ConnectionString string `json:"connection"`
	Database         string `json:"database"`
	Collection       string `json:"collection"`
}

type KafkaSettings struct {
	Address       string `json:"address"`
	ProducerTopic string `json:"producerTopic"`
	ConsumerTopic string `json:"consumerTopic"`
	ConsumerGroup string `json:"consumerGroup"`
}
