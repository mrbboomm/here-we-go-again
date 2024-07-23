package config

var KafkaTopics = []string{"tier", "user"}

type KafkaConnCfg struct {
	Url    string
	Topics []string
}
