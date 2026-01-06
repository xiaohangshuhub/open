package app

import (
	"time"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/go-workit/pkg/ddd"
	"github.com/xiaohangshuhub/open/internal/service1/domain/dic/status"
	"github.com/xiaohangshuhub/open/internal/service1/domain/err"
)

const (
	topicNameLengthMin = 2
	topicNameLengthMax = 20
)

// Topic 主题实体
type Topic struct {
	ddd.AggregateRoot[uuid.UUID]               // ID
	AppID                        uuid.UUID     // 所属应用ID
	Name                         string        // 名称
	DataStruct                   string        // 数据结构
	Desc                         string        // 描述
	Created                      time.Time     // 创建时间
	CreateBy                     uuid.UUID     // 创建人
	Updated                      *time.Time    // 修改时间
	UpdateBy                     *string       // 修改人
	Status                       status.Status // 状态
}

// NewTopic 创建主题实体
func NewTopic(appID uuid.UUID, name, desc, dataStruct string, createBy uuid.UUID) (*Topic, *err.Error) {

	topic := &Topic{
		AggregateRoot: ddd.NewAggregateRoot(uuid.New()),
		Status:        status.Enable,
		Created:       time.Now(),
	}

	if err := topic.SetAppID(appID); err != nil {
		return nil, err
	}
	if err := topic.SetName(name); err != nil {
		return nil, err
	}
	if err := topic.SetDesc(desc); err != nil {
		return nil, err
	}
	if err := topic.SetDataStruct(dataStruct); err != nil {
		return nil, err
	}
	if createBy == uuid.Nil {
		return nil, err.ErrCreateByEmpty
	}

	return topic, nil
}

// SetName 设置主题名称
func (t *Topic) SetName(name string) *err.Error {
	if name == "" {
		return err.ErrNameEmpty
	}
	if len := len(name); len > topicNameLengthMax || len < topicNameLengthMin {
		return err.ErrLengthOutOfRange
	}
	t.Name = name
	return nil
}

// SetAppID 设置主题所属应用ID
func (t *Topic) SetAppID(appID uuid.UUID) *err.Error {
	if appID == uuid.Nil {
		return err.ErrAppIDEmpty
	}
	t.AppID = appID
	return nil
}

// SetDesc 设置主题描述
func (t *Topic) SetDesc(desc string) *err.Error {
	if desc == "" {
		return err.ErrDescEmpty
	}
	t.Desc = desc
	return nil
}

// SetDataStruct 设置主题数据结构
func (t *Topic) SetDataStruct(dataStruct string) *err.Error {
	if dataStruct == "" {
		return err.ErrDataStructEmpty
	}
	t.DataStruct = dataStruct
	return nil
}
