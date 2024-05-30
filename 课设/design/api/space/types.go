package space

import (
	"design/domain/space"
	"gorm.io/gorm"
)

type CreateTrendRequest struct {
	UserId   uint      `json:"userId"`
	Detail   string    `json:"detail"`
	Praise   uint      `json:"praise"`
	Comments []Comment `json:"comments" gorm:"foreignKey:SpaceTrendsId"` //评论[](不计入表)
	SpaceId  uint      `json:"spaceId"`
}
type CreateSpaceResp struct {
	UserId uint
}

type Comment struct {
	gorm.Model
	UserId   uint //评论用户id
	User     User `gorm:"-"`
	Praise   uint
	Content  string //内容
	TrendsId uint   //空间动态表id
	ToUserId uint   //0
}
type User struct {
	Username string `gorm:"type:varchar(30)"`  //用户名
	Img      string `gorm:"type:varchar(500)"` //头像地址
}

func ToSpaceComment(comment Comment) space.Comment {
	return space.Comment{
		UserId:   comment.UserId,
		Praise:   comment.Praise,
		Content:  comment.Content,
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
		Comments: ToSpaceComments(trend.Comments),
		SpaceId:  trend.SpaceId,
	}
}

type CreateTrendResponse struct {
	SpaceId uint `json:"spaceId"`
}

type FindTrendRequest struct {
	UserId uint `json:"userId"`
}

type FindTrendResponse struct {
	Trends []space.SpaceTrends `json:"trends"`
}

type CreateCommentRequest struct {
	Comment space.Comment `json:"comment"`
	UserId  uint          `json:"userId"`
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
