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

func Pass(ctx *gin.Context) {
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
	case models.NEW: // 审批新建的社团
		var aps models.Applications
		via := utils.MD5(data.First + data.Second + strconv.Itoa(int(models.NEW)))
		// fmt.Println(via, data.First, data.Second, models.ADD)
		db.Where("via = ?", via).First(&aps)
		if aps.Via == "" {
			common.FailRes(ctx, nil, "错误许可")
			return
		}
		db.Where("via = ?", via).Model(&aps).Update("status", models.ACTIVE)
		var community models.Community
		community.Id = utils.GenIDByString()
		community.Name = aps.First
		community.Hostname = aps.Second
		community.Context = aps.Context
		community.Timestamp = utils.NowTimestampByString()
		community.Status = models.ACTIVE
		db.Insert(&community)
		common.SuccessRes(ctx, nil, "审批成功")
	default:
	}
}
