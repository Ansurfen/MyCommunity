package user

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"

	"github.com/gin-gonic/gin"
)

func Cancel(ctx *gin.Context) {
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	db := sql.GetDB("general")
	db.Delete(user)
	common.SuccessRes(ctx, nil, "注销成功")
}
