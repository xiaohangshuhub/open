package app

type ClinetType int8

const (
	WebApp     ClinetType = iota + 1 // web应用
	DesptopApp                       // 桌面应用
	MobileApp                        // 移动应用
	MiniApp                          // 小程序
	ClientApp                        // 客户端应用
)
