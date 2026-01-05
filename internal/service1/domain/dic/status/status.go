package status

type Status int8

const (
	Enable  Status = iota + 1 // 启用
	Disable                   // 禁用
	Review                    // 审核
)
