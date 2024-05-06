package usertoUser

import "errors"

// 用户自定义错误

var (
	ErrNotUser       = errors.New("您已删除该用户")
	ErrNotSend       = errors.New("发送异常")
	ErrNotRevocation = errors.New("撤回异常")
)
