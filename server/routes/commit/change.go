package commit

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChangeRight(ctx *gin.Context) {
	var data models.Applications
	if err := ctx.ShouldBindJSON(&data); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	if data.Type != models.CHANGE {
		common.FailRes(ctx, nil, "状态码校验失败")
		return
	}
	var user models.User
	if tmpl, b := ctx.Get("user"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "User don't exist")
		return
	} else {
		user = tmpl.(models.User)
	}
	if user.Right != models.ROOT {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "没有权限")
		return
	}
	db := sql.GetDB("general")
	db.Where("username = ?", data.First).Model(&models.User{}).Update("right", data.Status)
	common.SuccessRes(ctx, nil, "设置成功")
}
