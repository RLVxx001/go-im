package usertoUser

import "errors"

// 用户自定义错误

var (
	ErrNotUser       = errors.New("您已删除该用户")
	ErrNotSend       = errors.New("发送异常")
	ErrNotRevocation = errors.New("撤回异常")
	ErrNotDelete     = errors.New("删除异常")
	ErrNotFid        = errors.New("查找异常")
	ErrNotCreate     = errors.New("创建异常")
	ErrNotUpdate     = errors.New("修改异常")
)
