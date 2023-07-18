package http

import (
	"backend/internal/rank"
	"backend/pkg/myerr"
	"backend/pkg/result"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *RankHandler) GetRankList(ctx *gin.Context) {
	var getRankParam rank.GetRankListReq

	if err := ctx.ShouldBind(&getRankParam); err != nil {
		result.SendResult(ctx, result.Fail(myerr.ParamErr))
		return
	}

	page, _ := strconv.Atoi(getRankParam.Page)

	if page <= 0 {
		page = 1
	}

	resp, err := r.svc.GetRankList(ctx, page)
	if err != nil {
		result.SendResult(ctx, result.Fail(myerr.ParamErr))
		return
	}

	result.SendResult(ctx, result.Success("success", resp))
}
