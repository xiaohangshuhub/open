package app

import "errors"

type Error struct {
	Err  error
	Code int
}

func (e Error) Error() string {
	return e.Err.Error()
}

var (
	ErrUnknown                    = &Error{Code: 10000, Err: errors.New("unknown error")}                   // 未知错误
	ErrClientNameEmpty            = &Error{Code: 10001, Err: errors.New("client name is empty")}            // 客户端名称为空
	ErrClientNameLengthOutOfRange = &Error{Code: 10002, Err: errors.New("client name length Out of range")} // 客户端名称超出范围
	ErrCreateByEmpty              = &Error{Code: 10003, Err: errors.New("create_by is empty")}              // 创建人为空
)
