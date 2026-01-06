package err

import "errors"

type Error struct {
	Err  error
	Code int
}

func (e Error) Error() string {
	return e.Err.Error()
}

var (
	ErrUnknown          = &Error{Code: 10000, Err: errors.New("unknown error")}              // 未知错误
	ErrNameEmpty        = &Error{Code: 10001, Err: errors.New("name is empty")}              // 名称为空
	ErrLengthOutOfRange = &Error{Code: 10002, Err: errors.New("length Out of range")}        // 超出范围
	ErrCreateByEmpty    = &Error{Code: 10003, Err: errors.New("create_by is empty")}         // 创建人为空
	ErrDescEmpty        = &Error{Code: 10004, Err: errors.New("desc is empty")}              // 描述为空
	ErrAppIDEmpty       = &Error{Code: 10005, Err: errors.New("app id is empty")}            // 应用ID为空
	ErrDataStructEmpty  = &Error{Code: 10006, Err: errors.New("topic data struct is empty")} // 主题数据结构为空
)
