package user

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	username, key, password := ctx.PostForm("username"), ctx.PostForm("key"), ctx.PostForm("password")
	// key, password = utils.DecodeAESWithKey(utils.FrontKey, key), utils.DecodeAESWithKey(utils.FrontKey, password)
	if len(password) < 6 {
		common.FailRes(ctx, nil, "Password less than 6")
		return
	}
	var user *models.User
	if common.CheckPhone(key) {
		user = models.NewUserWithPhone(username, key, utils.HashPassword(password))
		user.Telephone = utils.EncodeAESWithKey(utils.BackedKey, user.Telephone)
	}
	if common.CheckEmail(key) {
		user = models.NewUserWithEmail(username, key, utils.HashPassword(password))
		user.Email = utils.EncodeAESWithKey(utils.BackedKey, user.Email)
	}
	if user.Telephone == "" && user.Email == "" {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Unknown key")
		return
	}
	db := sql.GetDB("general")
	db.Insert(&user)
	common.SuccessRes(ctx, nil, "Register successfully")
}
