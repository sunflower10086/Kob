package rank

import (
	"context"
)

type Service interface {
	GetRankList(context.Context, int) (*GetRankListResp, error)
}

type GetRankListReq struct {
	Page string `json:"page" binding:"required" form:"page" query:"page"`
}

type GetRankListResp struct {
	Users     *[]RankUser `json:"users,omitempty"`
	UserCount int         `json:"user_count,omitempty"`
}

type RankUser struct {
	Id       int32  `json:"id,omitempty"`
	Photo    string `json:"photo,omitempty"`
	UserName string `json:"username,omitempty"`
	Rating   *int32 `json:"rating,omitempty"`
	Number   int    `json:"number,omitempty"`
}
