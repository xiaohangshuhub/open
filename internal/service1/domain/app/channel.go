package app

import (
	"time"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/go-workit/pkg/ddd"
	"github.com/xiaohangshuhub/open/internal/service1/domain/err"
)

// 数据方向：相对于平台本身
type Direction int8

const (
	Inbound  Direction = iota + 1 // 外部 → 平台
	Outbound                      // 平台 → 外部
)

// 数据传输方式
type Transport int8

const (
	HTTP Transport = iota + 1
	RabbitMQ
	Kafka
	RocketMQ
	MQTT
)

// 认证类型
type AuthType int8

const (
	NoAuth AuthType = iota + 1
	APIKey
	OAuth2
)

// Channel 通道实体
type Channel struct {
	ddd.AggregateRoot[uuid.UUID]            // ID
	Name                         string     // 名称
	Direction                    Direction  // 数据方向
	Transport                    Transport  // 数据传输方式
	AuthType                     AuthType   // 认证类型
	Endpoint                     string     // 访问地址
	AccessKey                    *string    // 访问key
	SecretKey                    *string    // 密钥
	AppID                        string     // 所属应用ID
	Created                      time.Time  // 创建时间
	CreateBy                     string     // 创建人
	Updated                      *time.Time // 修改时间
	UpdateBy                     *string    // 修改人
}

// NewChannel 创建通道实体
func NewChannel(name string, direction Direction, transport Transport, appID string, createBy string) (*Channel, *err.Error) {
	channel := &Channel{
		AggregateRoot: ddd.NewAggregateRoot(uuid.New()),
		Direction:     direction,
		Transport:     transport,
		AuthType:      NoAuth,
		Created:       time.Now(),
	}

	if err := channel.SetName(name); err != nil {
		return nil, err
	}
	if err := channel.SetAppID(appID); err != nil {
		return nil, err
	}
	if createBy == "" {
		return nil, err.ErrCreateByEmpty
	}

	return channel, nil
}

// SetName 设置通道名称
func (c *Channel) SetName(name string) *err.Error {
	if name == "" {
		return err.ErrNameEmpty
	}
	c.Name = name
	return nil
}

// SetAppID 设置所属应用ID
func (c *Channel) SetAppID(appID string) *err.Error {
	if appID == "" {
		return err.ErrAppIDEmpty
	}
	c.AppID = appID
	return nil
}
