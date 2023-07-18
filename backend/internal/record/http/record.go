package http

import (
	"backend/internal/record"
	"backend/pkg/myerr"
	"backend/pkg/result"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *RecordHandler) GetRecordList(ctx *gin.Context) {
	var getListParam record.GetListReq

	if err := ctx.ShouldBind(&getListParam); err != nil {
		result.SendResult(ctx, result.Fail(myerr.ParamErr))
		return
	}

	page, _ := strconv.Atoi(getListParam.Page)

	if page <= 0 {
		page = 1
	}

	resp, err := r.svc.GetList(ctx, page)
	if err != nil {
		result.SendResult(ctx, result.Fail(err))
		return
	}

	result.SendResult(ctx, result.Success("success", resp))
}
