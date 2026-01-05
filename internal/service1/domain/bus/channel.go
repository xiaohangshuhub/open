package bus

type Channel struct {
	Producers ChannelType // 生产者类型
	Consumers ChannelType // 消费者类型
}
