package middlewares

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			common.CommonRes(ctx, http.StatusUnauthorized, nil, "权限不足")
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			common.CommonRes(ctx, http.StatusUnauthorized, nil, "权限不足")
			ctx.Abort()
			return
		}
		db := sql.GetDB("general")
		var user models.User
		db.Query(&user, claims.UID)
		if user.Id == 0 {
			common.CommonRes(ctx, http.StatusUnauthorized, nil, "权限不足")
			ctx.Abort()
			return
		}
		if user.Telephone != "" {
			user.Telephone = utils.DecodeAESWithKey(utils.BackedKey, user.Telephone)
		}
		if user.Email != "" {
			user.Email = utils.DecodeAESWithKey(utils.BackedKey, user.Email)
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
