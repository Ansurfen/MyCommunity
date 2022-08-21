package routes

import (
	"go-auth/common"
	"go-auth/db/sql"
	"go-auth/models"
	"go-auth/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ReleaseJWT(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	token, err := common.ReleaseToken(user)
	if err != nil {
		common.FailRes(ctx, nil, "系统错误")
		return
	}
	common.SuccessRes(ctx, gin.H{"token": token}, "Release token Successfully!")
}

func ParserToken(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		common.CommonRes(ctx, http.StatusUnauthorized, nil, "权限不足")
		return
	}
	tokenString = tokenString[7:]
	token, claims, err := common.ParseToken(tokenString)
	if err != nil || !token.Valid {
		common.CommonRes(ctx, http.StatusUnauthorized, nil, "权限不足")
		return
	}
	db := sql.GetDB("general")
	var user models.User
	db.Query(&user, claims.UID)
	if user.Id == 0 {
		common.CommonRes(ctx, http.StatusUnauthorized, nil, "权限不足")
		return
	}
	if user.Telephone != "" {
		user.Telephone = utils.DecodeAESWithKey(utils.BackedKey, user.Telephone)
	}
	if user.Email != "" {
		user.Email = utils.DecodeAESWithKey(utils.BackedKey, user.Email)
	}
	common.SuccessRes(ctx, gin.H{"user": user}, "Parser Pass!")
}
