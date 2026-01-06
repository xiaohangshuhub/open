package bus

// ChannelType 通道类型
type ChannelType int8

// 通道类型枚举
const (
	RabbitMQ ChannelType = iota + 1
	Kafka
	HTTP
)

// Channel 通道实体
type Channel struct {
	Producers ChannelType // 生产者类型
	Consumers ChannelType // 消费者类型
}
