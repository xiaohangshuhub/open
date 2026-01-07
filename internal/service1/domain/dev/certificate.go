package dev

import "github.com/xiaohangshuhub/go-workit/pkg/ddd"

// Certificate 开发者认证实体
type Certificate struct {
	ddd.AggregateRoot[string]        // ID
	DeveloperID               string // 开发者ID
	CertType                  int8   // 认证类型
	CertValue                 string // 认证值
	Company                   string // 企业
}
