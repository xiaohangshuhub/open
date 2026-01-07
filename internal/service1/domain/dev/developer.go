package dev

import "github.com/xiaohangshuhub/go-workit/pkg/ddd"

// Developer 开发者实体
type Developer struct {
	ddd.AggregateRoot[string]        // ID
	Name                      string // 名称
	Email                     string // 邮箱
	Company                   string // 公司
}
