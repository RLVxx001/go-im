package usertoUser

import (
	wsServer "design/api/ws"
	"design/domain/user"
	"design/domain/usertoUser"
	"design/utils/api_helper"
	"design/utils/webSocketDecoded"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Controller struct {
	service     *usertoUser.Service
	userService *user.Service
}

// 实例化
func NewController(service *usertoUser.Service, userService *user.Service) *Controller {
	return &Controller{service: service, userService: userService}
}

// 创建用户-用户链接
func (c *Controller) Create() {

	for {
		var req UserRequest
		body := <-wsServer.UserChan
		fmt.Printf("%v\n", body)
		//校验用户id是否正确
		req.UserOwner = body.Owner
		req.UserTarget = body.Target
		fmt.Printf("%v\n", req)

		if _, err := c.userService.GetById(req.UserOwner); err != nil {
			log.Println(err)
			continue
		}

		if _, err := c.userService.GetById(req.UserTarget); err != nil {
			log.Println(err)
			continue
		}
		//创建链接
		user1, err := c.service.Create(usertoUser.NewUsertoUser(req.UserOwner, req.UserTarget, body.Remarks))
		if err != nil {
			log.Println(err)
			continue
		}
		response1 := ToUserResponse(user1)
		us, err1 := c.userService.GetById(response1.UserTarget)
		if err1 == nil {
			response1.ToUser.Username = us.Username
			response1.ToUser.Account = us.Account
			response1.ToUser.Img = us.Img
		}
		var user2 *usertoUser.UsertoUser = nil
		user2, err = c.service.Create(usertoUser.NewUsertoUser(req.UserTarget, req.UserOwner, body.Remarks1))
		if err != nil {
			log.Println(err)
			continue
		}
		response2 := ToUserResponse(user2)
		us, err1 = c.userService.GetById(response1.UserTarget)
		if err1 == nil {
			response2.ToUser.Username = us.Username
			response2.ToUser.Account = us.Account
			response2.ToUser.Img = us.Img
		}
		c.service.Send(user1, body.Remarks1)
		c.service.Send(user2, body.Remarks)
		wsServer.Broadcast <- wsServer.NewW(user1.UserOwner, response1, body.Event) //发送者
		wsServer.Broadcast <- wsServer.NewW(user2.UserOwner, response2, body.Event) //发送者

	}
}

// 发送消息
func (c *Controller) Send(ws *websocket.Conn, mp map[string]interface{}, userid uint) {

	fmt.Println("发送链接中--------")

	fmt.Println("验证成功：userid:", userid)

	var req UserRequest
	err := webSocketDecoded.DecodedMap(mp, &req)
	fmt.Printf("%v\n%v\n", req, userid)
	if err != nil {
		log.Printf("error: %v\n", err)
		err := api_helper.WsError(ws, err, "auth")
		if err != nil {
			return
		}
	}

	req.UserOwner = userid
	if _, err := c.userService.GetById(req.UserOwner); err != nil {
		_ = api_helper.WsError(ws, api_helper.ErrInvalidToken, "token")
		return
	}

	utou := ToUsertoUser(req)
	//fmt.Printf("%v\n", utou)
	m, m1, err := c.service.Send(utou, req.Message)
	if err != nil {
		err1 := api_helper.WsError(ws, err, "err")
		if err1 != nil {
			log.Printf("error: %v", err1)
		}
		return
	}

	wsServer.Broadcast <- wsServer.NewW(req.UserOwner, UserMessage{
		Message:      m.Message,
		UsertoUserId: m.UsertoUserId,
		Key:          m.Key,
		User:         req.UserOwner,
		UserOwner:    req.UserOwner,
		CreatedAt:    m.CreatedAt,
	}, mp["event"].(string)) //发送者
	// Send the newly received message to the broadcast channel
	wsServer.Broadcast <- wsServer.NewW(req.UserTarget, UserMessage{
		Message:      m1.Message,
		UsertoUserId: m1.UsertoUserId,
		Key:          m1.Key,
		User:         req.UserTarget,
		UserOwner:    req.UserOwner,
		CreatedAt:    m1.CreatedAt,
	}, mp["event"].(string)) //送达者

}

// 修改用户-用户信息
func (c *Controller) Update(g *gin.Context) {

	var req UserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	req.UserOwner = api_helper.GetUserId(g)
	utou := ToUsertoUser(req)

	if err := c.service.Update(utou); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, nil)
}

// 撤回消息
func (c *Controller) Revocation(ws *websocket.Conn, mp map[string]interface{}, userid uint) {
	fmt.Println("验证成功：userid:", userid)

	var req UserRequest
	err := webSocketDecoded.DecodedMap(mp, &req)
	//fmt.Printf("req:%v\n", req)
	if err != nil {
		err := api_helper.WsError(ws, api_helper.ErrInvalidBody, "auth")
		if err != nil {
			return
		}
		return
	}
	if len(req.UserMessages) == 0 {
		err := api_helper.WsError(ws, api_helper.ErrInvalidBody, "token")
		if err != nil {
			return
		}
		return
	}
	req.UserOwner = userid
	if _, err := c.userService.GetById(req.UserOwner); err != nil {
		_ = api_helper.WsError(ws, api_helper.ErrInvalidToken, "err")
		return
	}
	utou := ToUsertoUser(req)
	fidutou, err := c.service.Fid(utou.UserOwner, utou.UserTarget)
	if err != nil {
		err := api_helper.WsError(ws, err, "err")
		if err != nil {
			return
		}
		return
	}
	utou.ID = fidutou.ID
	if err := c.service.Revocation(utou); err != nil {
		err := api_helper.WsError(ws, err, "err")
		if err != nil {
			return
		}
		return
	}

	for _, j := range utou.UserMessages {
		wsServer.Broadcast <- wsServer.NewW(req.UserOwner, UserMessage{
			UsertoUserId: utou.ID,
			Key:          j.Key,
		}, mp["event"].(string)) //发送者
		// Send the newly received message to the broadcast channel
		wsServer.Broadcast <- wsServer.NewW(req.UserTarget, UserMessage{
			UsertoUserId: utou.ID,
			Key:          j.Key,
		}, mp["event"].(string)) //送达者
		break //暂时先处理一个
	}

}

// 用户单方面删除消息
func (c *Controller) DeleteMessage(g *gin.Context) {
	var req UserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	if len(req.UserMessages) == 0 {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	req.UserOwner = api_helper.GetUserId(g)
	utou := ToUsertoUser(req)
	if err := c.service.DeleteMessage(utou); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, nil)
}

// 用户单方面删除所有消息
func (c *Controller) DeleteMessages(g *gin.Context) {
	var req UserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	req.UserOwner = api_helper.GetUserId(g)
	utou := ToUsertoUser(req)
	if err := c.service.DeleteMessages(utou); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, nil)
}

// 删除好友
func (c *Controller) DeleteUser(g *gin.Context) {
	var req UserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	req.UserOwner = api_helper.GetUserId(g)

	fid, err := c.service.Fid(req.UserOwner, req.UserTarget)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	fid.IsDeleted = true
	err = c.service.DeleteUser(fid)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, nil)
}

// 查找好友信息
func (c *Controller) Fids(g *gin.Context) {
	userid := api_helper.GetUserId(g)
	fids, err := c.service.Fids(userid)

	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	var userResponses []UserResponse
	for _, j := range fids {
		response := ToUserResponse(&j)
		us, err := c.userService.GetById(j.UserTarget)
		if err == nil {
			j.ToUser = us
		}
		response.ToUser.Username = j.ToUser.Username
		response.ToUser.Account = j.ToUser.Account
		response.ToUser.UserId = j.ToUser.ID
		response.ToUser.Img = j.ToUser.Img
		userResponses = append(userResponses, response)
	}

	g.JSON(http.StatusOK, userResponses)
}

// 查找好友信息
func (c *Controller) Fid(g *gin.Context) {
	var req UserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	userid := api_helper.GetUserId(g)
	fid, err := c.service.FidMessage(usertoUser.NewUsertoUser(userid, req.UserTarget, ""))
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	response := ToUserResponse(fid)
	user, err := c.userService.GetById(fid.UserTarget)
	if err == nil {
		response.ToUser.Username = user.Username
		response.ToUser.Account = user.Account
		response.ToUser.UserId = user.ID
		response.ToUser.Img = user.Img
	}
	g.JSON(http.StatusOK, response)
}

// 查看消息
func (c *Controller) Read(g *gin.Context) {
	var req UserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	req.UserOwner = api_helper.GetUserId(g)
	fid, err := c.service.Fid(req.UserOwner, req.UserTarget)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	c.service.ReadMessage(fid.ID)
	g.JSON(http.StatusOK, nil)
}
