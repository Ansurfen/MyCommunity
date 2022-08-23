package middlewares

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckUniqueUsername() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		db := sql.GetDB("general")
		var user models.User
		db.Where("username = ?", username).First(&user)
		if user.Id != 0 {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "User exist")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func CheckUniqueKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// key := utils.DecodeAESWithKey(utils.FrontKey, ctx.PostForm("key"))
		key := ctx.PostForm("key")
		isValid := false
		db := sql.GetDB("general")
		var user models.User
		if common.CheckEmail(key) {
			isValid = true
			db.Where("email = ?", key).First(&user)
			if user.Id != 0 {
				common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Email exist")
				ctx.Abort()
				return
			}
			ctx.Next()
		}
		if common.CheckPhone(key) {
			isValid = true
			db.Where("telephone = ?", key).First(&user)
			if user.Id != 0 {
				common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Telephone exist")
				ctx.Abort()
				return
			}
			ctx.Next()
		}
		if !isValid {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Unknown key")
			ctx.Abort()
			return
		}
	}
}

func CheckUniqueTelephone() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newPhone := ctx.PostForm("newphone")
		newPhone = utils.DecodeAESWithKey(utils.FrontKey, newPhone)
		if common.CheckPhone(newPhone) {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "手机号不正确")
			ctx.Abort()
			return
		}
		var user models.User
		db := sql.GetDB("general")
		db.Where("telephone = ?", newPhone).First(&user)
		if user.Id != 0 {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "手机号已存在")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func CheckUniqueEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newEmail := ctx.PostForm("newemail")
		newEmail = utils.DecodeAESWithKey(utils.FrontKey, newEmail)
		if common.CheckEmail(newEmail) {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "邮箱不正确")
			ctx.Abort()
			return
		}
		var user models.User
		db := sql.GetDB("general")
		db.Where("email = ?", newEmail).First(&user)
		if user.Id != 0 {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "邮箱已存在")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
