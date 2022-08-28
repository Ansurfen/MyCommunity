package commit

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Reject(ctx *gin.Context) {
	var data models.Applications
	if tmpl, b := ctx.Get("application"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "申请不存在")
		return
	} else {
		data = tmpl.(models.Applications)
	}
	if data.Type != models.PASS {
		common.FailRes(ctx, nil, "状态码校验失败")
		return
	}
	db := sql.GetDB("general")
	switch data.Status {
	case models.NEW: // 审批新建的社团
		var aps models.Applications
		via := utils.MD5(data.First + data.Second + strconv.Itoa(int(models.NEW)))
		db.Where("via = ?", via).First(&aps)
		if aps.Via == "" {
			common.FailRes(ctx, nil, "错误许可")
			return
		}
		db.Where("via = ?", via).Model(&aps).Update("status", models.BANNED)
		common.SuccessRes(ctx, nil, "审批成功")
	case models.ADD:
		print("拒绝申请")
	default:
	}
}
