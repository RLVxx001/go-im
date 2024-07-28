package userImg

import "errors"

// 用户自定义错误

var (
	ErrFind   = errors.New("查找失败")
	ErrDelete = errors.New("图片删除异常")
)
