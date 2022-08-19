package commit

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Reject(ctx *gin.Context) {
	var data models.Applications
	if err := ctx.ShouldBindJSON(&data); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	if data.Type != models.PASS {
		common.FailRes(ctx, nil, "状态码校验失败")
		return
	}
	db := sql.GetDB("general")
	switch data.Status {
	case models.ADD: // 审批新建的社团
		var aps models.Applications
		via := utils.MD5(data.First + data.Second + strconv.Itoa(int(models.ADD)))
		fmt.Println(via, data.First, data.Second, models.ADD)
		db.Where("via = ?", via).First(&aps)
		if aps.Via == "" {
			common.FailRes(ctx, nil, "错误许可")
			return
		}
		db.Where("via = ?", via).Model(&aps).Update("status", models.BANNED)
		common.SuccessRes(ctx, nil, "审批成功")
	default:
	}
}
