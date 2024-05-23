package webSocketDecoded

import (
	"design/utils/api_helper"
	"design/utils/jwt"
	"design/utils/pagination"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

// 专门负责解析ws请求中的token与event和其他请求体 并且返回
func Decoded(ws *websocket.Conn, userid *uint, event *string) (error, string, map[string]interface{}) {

	mp := make(map[string]interface{})
	err := ws.ReadJSON(&mp)
	if err != nil {
		return api_helper.ErrInvalidBody, "", nil
	}
	for i, j := range mp {
		fmt.Printf("%v  %v\n", i, j)
	}
	token, ok := mp["token"].(string)
	if !ok {
		return api_helper.ErrInvalidBody, "", nil
	}

	*event, ok = mp["event"].(string)
	if !ok {
		return api_helper.ErrInvalidBody, "", nil
	}

	id, err := jwt.Decoded(token)
	if err != nil {
		return api_helper.ErrInvalidToken, "token", nil
	}
	*userid = uint(pagination.ParseInt(id, 0))

	return nil, "", mp
}

// 解析map里面值
func DecodedMap(mp map[string]interface{}, req interface{}) error {
	err := mapstructure.Decode(mp, req) //map转指针结构体
	if err != nil {
		return api_helper.ErrInvalidBody
	}
	fmt.Printf("%v\n", req)
	return nil
}
