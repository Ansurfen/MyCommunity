package community

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	NOLOGIN = iota
	NOADD
	WAITING
	ADDED
	ADMIN
)

func Info(ctx *gin.Context) {
	cname := ctx.PostForm("cname")
	db := sql.GetDB("general")
	var community models.Community
	db.Where("name = ?", cname).Find(&community)
	res, err := json.Marshal(&community)
	if err != nil {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "系统错误")
		return
	}
	// 需要判断是否在社团里面了，是就不允许提交
	status := NOLOGIN
	if has, b := ctx.Get("has"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "系统错误")
		return
	} else {
		if has.(bool) == true {
			status = NOADD
			var user *models.User
			if user = common.GetUser(ctx); user == nil {
				status = NOLOGIN
			} else {
				if user.Username == community.Hostname {
					status = ADMIN
				}
				var admins models.Admins
				if err := json.Unmarshal([]byte(community.Admins), &admins); err == nil {
					for _, admin := range admins {
						if user.Username == admin {
							status = ADMIN
						}
					}
				}
				var memebers models.Members
				if err := json.Unmarshal([]byte(community.Members), &memebers); err == nil {
					for _, member := range memebers {
						if user.Username == member {
							status = ADDED
						}
					}
				}
				if status == NOADD {
					count := 0
					db.Where("first = ? and second = ?", user.Username, community.Name).Model(&models.Applications{}).Count(&count)
					if count > 0 {
						status = WAITING
					}
				}
			}
		} else {
			status = NOLOGIN
		}
	}
	common.SuccessRes(ctx, gin.H{"data": string(res), "status": status}, "获取数据成功")
}

func InfoApplication(ctx *gin.Context) {
	// cname := ctx.PostForm("cname")
	// var community models.Community
	// db := sql.GetDB("general")
	// db.Where("name = ?", cname).First(&community)
	// if community.Name != "" {
	// 	common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community name exist")
	// 	return
	// }
	// var applications []models.Application
	// db.Where("to = ?", cname).Find(&applications)
	// if data, err := json.Marshal(applications); err != nil {
	// 	common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "System error")
	// 	return
	// } else {
	// 	common.SuccessRes(ctx, gin.H{
	// 		"data": data,
	// 	}, "")
	// }
}
