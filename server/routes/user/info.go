package user

import (
	"MyCommunity/common"
	"MyCommunity/models"

	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context) {
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	common.SuccessRes(ctx, gin.H{
		"username":  user.Username,
		"alias":     user.Alias,
		"telephone": user.Telephone,
		"email":     user.Email,
		"school":    user.School,
		"right":     user.Right,
		"id":        user.Id,
		"profile":   user.Profile,
	}, "信息获取成功")
}
