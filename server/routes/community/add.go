package community

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
	first -> from
	second -> to
	type -> add
*/
// 申请加入社团
func Add(ctx *gin.Context) {
	var data models.Applications
	if err := ctx.ShouldBindJSON(&data); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	if data.Type != models.ADD {
		common.FailRes(ctx, nil, "状态码校验失败")
		return
	}
	db := sql.GetDB("general")
	var aps models.Applications
	data.Via = utils.MD5(data.First + data.Second + strconv.Itoa(int(data.Type)))
	db.Where("via = ?", data.Via).First(&aps)
	if aps.Via != "" {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "请勿重复提交申请")
		return
	}
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	if user.Username != data.First {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "User's check err")
		return
	}
	data.Status = models.WAITING
	db.Insert(&data)
	common.SuccessRes(ctx, nil, "提交成功，等待审核")
}
