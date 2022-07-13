package user

import (
	"MyCommunity/cache"
	"MyCommunity/common"
	"MyCommunity/db/nosql"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Login(ctx *gin.Context) {
	var data models.User
	if err := ctx.ShouldBindJSON(&data); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	if len(data.Username) == 0 {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "User don't empty")
		return
	}
	var user models.User
	db := sql.GetDB("general")
	db.Where("username = ?", data.Username).First(&user)
	if user.Id == 0 {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "User don't exist")
		return
	}
	if err := utils.EqualHashPassword(data.Password, user.Password); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, nil, "Password error")
		return
	}
	var token string
	if tmpl, err := cache.Cache.Get([]byte(user.Username)); err == nil {
		token = parserToken(string(tmpl))
	}
	if token == "" {
		if tmpl, err := nosql.GetRedis().Get(user.Username).Result(); err == nil {
			token = parserToken(tmpl)
		}
	}
	if token == "" {
		var err error
		token, err = common.ReleaseToken(user)
		if err != nil {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "System error")
			logrus.Error("Token generate error: %v", err)
			return
		}
	}
	cache.Cache.Set([]byte(user.Username), []byte(token), cache.Week)
	nosql.GetRedis().Set(user.Username, token, time.Hour*24*7)
	common.SuccessRes(ctx, gin.H{"token": token}, "Login Successfully!")
}

func parserToken(tmpl string) string {
	_, claims, err := common.ParseToken(string(tmpl))
	if !claims.IsTimeout() && err != nil {
		return tmpl
	}
	return ""
}