package group

import "errors"

// 用户自定义错误

var (
	ErrGroupId       = errors.New("该群号已存在")
	ErrNotGroupId    = errors.New("该群不存在")
	ErrNotSend       = errors.New("发送异常")
	ErrNotCreateUser = errors.New("该用户已存在群中")
	ErrNotDelete     = errors.New("删除异常")
	ErrNotGroupUser  = errors.New("群中没有该用户")
	ErrCreate        = errors.New("创建异常")
	ErrNotUpdate     = errors.New("对不起您的权限不够")
	ErrUpdate        = errors.New("修改异常")
	ErrGag           = errors.New("已被禁言")
	ErrRevocation    = errors.New("撤回异常")
	ErrFid           = errors.New("查找群聊异常")
)
