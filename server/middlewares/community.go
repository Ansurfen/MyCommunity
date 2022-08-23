package middlewares

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckCommunityRight(right int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var post models.Post
		if err := ctx.ShouldBindJSON(&post); err != nil {
			common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
			return
		}
		var community models.Community
		db := sql.GetDB("general")
		db.Where("id = ?", post.Id).First(&community)
		if community.Name == "" {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community don't exist")
			return
		}
		var user *models.User
		if user = common.GetUser(ctx); user == nil {
			return
		}
		switch right {
		// HOST
		case models.ROOT:
			if community.Hostname != user.Username {
				common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "没有权限")
				return
			}
		// Admins
		case models.ADMIN:
			pass := false
			if community.Hostname == user.Username {
				pass = true
			}
			if community.Admins != "" && !pass {
				var admins models.Admins
				if err := json.Unmarshal([]byte(community.Admins), &admins); err != nil {
					common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
					return
				}
				for _, admin := range admins {
					if admin == user.Username {
						pass = true
					}
				}
			}
			if !pass {
				common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "没有权限")
				return
			}
		default:
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "没有权限")
			return
		}
		ctx.Set("post", post)
		ctx.Set("community", community)
		ctx.Next()
	}
}

func CheckCommunityMember() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var post models.Post
		if err := ctx.ShouldBindJSON(&post); err != nil {
			common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
			return
		}
		var community models.Community
		db := sql.GetDB("general")
		db.Where("id = ?", post.Id).First(&community)
		if community.Name == "" {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community don't exist")
			return
		}
		var user *models.User
		if user = common.GetUser(ctx); user == nil {
			return
		}
		exist := false
		if community.Hostname == user.Username {
			exist = true
		}
		if community.Admins != "" && !exist {
			var admins models.Admins
			if err := json.Unmarshal([]byte(community.Admins), &admins); err != nil {
				common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
				return
			}
			for _, admin := range admins {
				if admin == user.Username {
					exist = true
				}
			}
		}
		if community.Members != "" && !exist {
			var memebers models.Members
			if err := json.Unmarshal([]byte(community.Members), &memebers); err != nil {
				common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
				return
			}
			for _, memeber := range memebers {
				if memeber == user.Username {
					exist = true
				}
			}
		}
		if !exist {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "User don't join community")
			return
		}
		ctx.Set("post", post)
		ctx.Set("community", community)
		ctx.Next()
	}
}
