package bus

type ChannelType int8

const (
	RabbitMQ ChannelType = iota + 1
	Kafka
	HTTP
)
