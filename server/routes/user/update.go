package user

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"

	"github.com/gin-gonic/gin"
)

const savePath string = "./images/user/"

func Update(ctx *gin.Context) {
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	if file, err := ctx.FormFile("file"); err != nil {
		common.FailRes(ctx, nil, "上传失败")
	} else {
		ctx.SaveUploadedFile(file, savePath+user.Username+".png")
	}
	db := sql.GetDB("general")
	user.Profile=1
	db.Where("username = ?", user.Username).Model(&models.User{}).Update("profile", user.Profile)
	common.SuccessRes(ctx, nil, "上传成功")
}
