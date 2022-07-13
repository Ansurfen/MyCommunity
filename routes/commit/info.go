package commit

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context) {
	var communities []models.Community
	db := sql.GetDB("general")
	db.Where("status = ?", models.WAITING).Find(&communities)
	if len(communities) == 0 {
		common.SuccessRes(ctx, nil, "没有待审核的请求")
		return
	}
	if data, err := json.Marshal(communities); err != nil {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "System error")
		return
	} else {
		common.SuccessRes(ctx, gin.H{
			"data": data,
		}, "请求成功")
	}
}
