package middlewares

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if len(tokenString) == 0 {
			common.CommonRes(ctx, http.StatusUnauthorized, nil, "权限不足")
			ctx.Abort()
			return
		}
		if strings.HasPrefix(tokenString, "{") {
			tokenString = tokenString[1:]
			str := "{"
			i := 0
			for i = 0; tokenString[i] != '}'; i++ {
				str += string(tokenString[i])
			}
			str += "}"
			var header models.AuthHeader
			if err := json.Unmarshal([]byte(str), &header); err != nil {
				common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
				return
			}
			ctx.Set("header", header)
			tokenString = tokenString[i+1:]
		}
		if !strings.HasPrefix(tokenString, "Bearer ") {
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

func AuthJWTWithNull() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.Set("has", false)
			ctx.Next()
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.Set("has", false)
			ctx.Next()
			ctx.Abort()
			return
		}
		db := sql.GetDB("general")
		var user models.User
		db.Query(&user, claims.UID)
		if user.Id == 0 {
			ctx.Set("has", false)
			ctx.Next()
			ctx.Abort()
			return
		}
		if user.Telephone != "" {
			user.Telephone = utils.DecodeAESWithKey(utils.BackedKey, user.Telephone)
		}
		if user.Email != "" {
			user.Email = utils.DecodeAESWithKey(utils.BackedKey, user.Email)
		}
		ctx.Set("has", true)
		ctx.Set("user", user)
		ctx.Next()
	}
}
