package space

import "design/domain/space"

type CreateTrendRequest struct {
	Trend space.SpaceTrends `json:"trend"`
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
