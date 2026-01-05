package bus

import (
	"github.com/google/uuid"
	"github.com/xiaohangshuhub/go-workit/pkg/ddd"
)

type Topic struct {
	ddd.AggregateRoot[uuid.UUID]        // ID
	Topic                        string // 主题
	Name                         string // 名称

}
