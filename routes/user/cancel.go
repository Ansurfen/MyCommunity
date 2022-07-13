package user

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cancel(ctx *gin.Context) {
	var user models.User
	if tmpl, b := ctx.Get("user"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "用户不存在")
		return
	} else {
		user = tmpl.(models.User)
	}
	db := sql.GetDB("general")
	db.Delete(user)
	common.SuccessRes(ctx, nil, "注销成功")
}
