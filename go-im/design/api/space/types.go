package space

import (
	"design/domain/space"
	"design/domain/user"
	"time"
)

type CreateTrendRequest struct {
	UserId uint   `json:"userId"`
	Detail string `json:"detail"`
	Praise uint   `json:"praise"`
}
type CreateSpaceResp struct {
	UserId uint `json:"userId"`
}

type Comment struct {
	UserId   uint      `json:"userId"` //评论用户id
	User     user.User `json:"user" gorm:"-"`
	Praise   uint      `json:"praise"`
	Detail   string    `json:"detail"`   //内容
	TrendsId uint      `json:"trendsId"` //空间动态表id
	ToUserId uint      `json:"toUserId"` //0
}

type FindCommentRequest struct {
	TrendId uint `json:"trendId"`
}

type User struct {
	Username string `gorm:"type:varchar(30)"`  //用户名
	Img      string `gorm:"type:varchar(500)"` //头像地址
}

type FindMessageResp struct {
	ID      uint      `json:"Id"`
	User    user.User `json:"user" gorm:"-"`
	SpaceId uint      `json:"spaceId"`
	UserId  uint      `json:"userId"`
	Detail  string    `json:"detail"`
	Tim     time.Time `json:"tim"`
}

type DelMessageRequest struct {
	MessageId uint `json:"messageId"`
}

type CreateMessageRequest struct {
	UserId uint   `json:"userId"`
	Detail string `json:"detail"`
}

func ToFindMessageResp(message space.Message) FindMessageResp {
	return FindMessageResp{
		ID:     message.ID,
		User:   message.User,
		UserId: message.UserId,
		Tim:    message.CreatedAt,
		Detail: message.Detail,
	}
}

func ToFindMessageResps(message []space.Message) []FindMessageResp {
	var messages []FindMessageResp
	for i := 0; i < len(message); i++ {
		messages = append(messages, ToFindMessageResp(message[i]))
	}
	return messages
}

type FindMessageRequest struct {
	UserId uint `json:"userId"`
}

func ToSpaceComment(comment Comment) space.Comment {
	return space.Comment{
		UserId:   comment.UserId,
		Praise:   comment.Praise,
		Content:  comment.Detail,
		TrendsId: comment.TrendsId,
		ToUserId: comment.ToUserId,
	}
}

func ToSpaceComments(comment []Comment) []space.Comment {
	var comments []space.Comment
	for i := 0; i < len(comment); i++ {
		comments = append(comments, ToSpaceComment(comment[i]))
	}
	return comments
}

func ToSpaceTrend(trend CreateTrendRequest) space.SpaceTrends {
	return space.SpaceTrends{
		UserId:   trend.UserId,
		Detail:   trend.Detail,
		Praise:   trend.Praise,
		Comments: []space.Comment{},
		SpaceId:  0,
	}
}

type CreateTrendResponse struct {
	SpaceId uint `json:"spaceId"`
}

type FindTrendsRequest struct {
	UserId uint `json:"userId"`
}

type FindTrendResponse struct {
	UserId   uint      `json:"userId"`
	User     user.User `json:"user"  gorm:"-"`
	Detail   string    `json:"detail"`
	TrendId  uint      `json:"trendId"`
	Tim      time.Time `json:"tim"`
	Praise   uint      `json:"praise"`
	Comments []Comment `json:"comments"` //评论[](不计入表)
	SpaceId  uint      `json:"spaceId"`
}

type FindTrendRequest struct {
	TrendId uint `json:"trendId"`
}

func ToComment(comment space.Comment) Comment {
	return Comment{
		UserId:   comment.UserId,
		User:     comment.User,
		Praise:   comment.Praise,
		Detail:   comment.Content,
		TrendsId: comment.TrendsId,
		ToUserId: comment.ToUserId,
	}
}

func ToComments(comment []space.Comment) []Comment {
	var comments []Comment
	for i := 0; i < len(comment); i++ {
		tmp := comment[i]
		comments = append(comments, ToComment(tmp))
	}
	return comments
}

func ToFindTrendsResp(trends space.SpaceTrends) FindTrendResponse {
	return FindTrendResponse{
		UserId:   trends.UserId,
		User:     trends.User,
		Detail:   trends.Detail,
		Praise:   trends.Praise,
		Tim:      trends.CreatedAt,
		TrendId:  trends.ID,
		Comments: ToComments(trends.Comments),
		SpaceId:  trends.SpaceId,
	}
}

func ToFindTrendsResps(trend []space.SpaceTrends) []FindTrendResponse {
	var trends []FindTrendResponse
	for i := 0; i < len(trend); i++ {
		tmp := trend[i]
		trends = append(trends, ToFindTrendsResp(tmp))
	}
	return trends
}

type CreateCommentRequest struct {
	Detail  string `json:"detail"`
	UserId  uint   `json:"userId"`
	TrendId uint   `json:"trendId"`
}

type CreateCommentResponse struct {
	Comment space.Comment `json:"comment"`
}

type DeleteTrendRequest struct {
	Trend space.SpaceTrends `json:"trend"`
}

type DeleteTrendResponse struct {
	Trend space.SpaceTrends `json:"trend"`
}
