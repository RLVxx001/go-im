package userApplication

import "errors"

// 用户自定义错误

var (
	ErrApplication = errors.New("申请失败")
	ErrNotFid      = errors.New("查找异常")
	ErrNotUpdate   = errors.New("拒绝异常")
	ErrUser        = errors.New("您已添加对方为好友")
	ErrAccept      = errors.New("接受异常")
)
