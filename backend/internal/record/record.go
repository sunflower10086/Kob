package record

import (
	"backend/internal/models"
	"context"
)

type Service interface {
	GetList(context.Context, int) (*GetListResp, error)
}

type GetListReq struct {
	Page string `json:"page" binding:"required" form:"page" query:"page"`
}

type GetListResp struct {
	Records      *[]Record `json:"records"`
	RecordsCount int       `json:"records_count"`
}

type Record struct {
	APhoto    string         `json:"a_photo,omitempty"`
	AUserName string         `json:"a_username,omitempty"`
	BPhoto    string         `json:"b_photo,omitempty"`
	BUserName string         `json:"b_username,omitempty"`
	Result    string         `json:"result,omitempty"`
	Record    *models.Record `json:"record,omitempty"`
}
