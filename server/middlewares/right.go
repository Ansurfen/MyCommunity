package middlewares

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRight(demand int8) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.Applications
		if err := ctx.ShouldBindJSON(&data); err != nil {
			common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "系统错误: 未能解析JSON")
			return
		}
		var user *models.User
		if user = common.GetUser(ctx); user == nil {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "用户不存在")
			return
		}
		db := sql.GetDB("general")
		switch data.Status {
		case models.NEW, models.CHANGE, models.USERINFO:
			if user.Right > demand {
				common.FailRes(ctx, nil, "你没有权限执行此操作")
				ctx.Abort()
				return
			}
		case models.ADD:
			var community models.Community
			db.Where("name = ?", data.Second).Model(&models.Community{}).First(&community)
			flag := false
			if user.Username == community.Hostname {
				flag = true
			}
			var admins models.Admins
			if len(community.Admins) == 0 {
				community.Admins = "[]"
			}
			if err := json.Unmarshal([]byte(community.Admins), &admins); err != nil {
				common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "系统错误: 未能解析JSON")
				return
			}
			for _, admin := range admins {
				if user.Username == admin {
					flag = true
				}
			}
			if !flag {
				common.FailRes(ctx, nil, "你没有权限执行此操作")
				ctx.Abort()
				return
			}
			ctx.Set("community", community)
		default:
			common.FailRes(ctx, nil, "未知类型")
			ctx.Abort()
			return
		}
		ctx.Set("application", data)
		ctx.Next()
	}
}
