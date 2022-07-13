package user

import (
	"MyCommunity/common"
	"MyCommunity/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context) {
	var user models.User
	if tmpl, b := ctx.Get("user"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "User don't exist")
		return
	} else {
		user = tmpl.(models.User)
	}
	common.SuccessRes(ctx, gin.H{
		"username":  user.Username,
		"alias":     user.Alias,
		"telephone": user.Telephone,
		"email":     user.Email,
		"school":    user.School,
		"right":     user.Right,
	}, "信息获取成功")
}
