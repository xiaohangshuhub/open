package app

import (
	"time"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/go-workit/pkg/ddd"
	"github.com/xiaohangshuhub/open/internal/service1/domain/dic/status"
)

const (
	nameLengthMin = 2
	nameLengthMax = 10
)

type Client struct {
	ddd.AggregateRoot[uuid.UUID]               // ID
	Name                         string        // 名称
	Key                          string        // app key
	Security                     string        // app security
	ClinetType                   ClinetType    // 客户端类型
	Status                       status.Status // 状态
	Created                      time.Time     // 创建时间
	CreateBy                     uuid.UUID     // 创建人
	Updated                      *time.Time    // 修改时间
	UpdateBy                     *string       // 修改人
}

func NewClient(name string, createBy uuid.UUID, clientType ClinetType) (*Client, *Error) {

	client := Client{
		AggregateRoot: ddd.NewAggregateRoot(uuid.New()),
		Key:           uuid.NewString(),
		Security:      uuid.NewString(),
		Status:        status.Review,
		Created:       time.Now(),
		ClinetType:    clientType,
	}

	if err := client.SetName(name); err != nil {
		return nil, err
	}

	if createBy == uuid.Nil {
		return nil, ErrCreateByEmpty
	}

	client.CreateBy = createBy

	return &client, nil
}

// SetName 设置客户端的名称
func (c *Client) SetName(name string) *Error {
	if name == "" {
		return ErrClientNameEmpty
	}
	if len := len(name); len > nameLengthMax || len < nameLengthMin {
		return ErrClientNameLengthOutOfRange
	}
	c.Name = name
	return nil
}
