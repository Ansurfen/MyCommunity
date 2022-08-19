package community

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cancel(ctx *gin.Context) {
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	var community models.Community
	if tmpl, b := ctx.Get("community"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community don't exist")
		return
	} else {
		community = tmpl.(models.Community)
	}
	if community.Hostname == user.Username {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Host无法退出社团，只能解散")
		return
	}
	db := sql.GetDB("general")
	if len(community.Admins) != 0 {
		var admins models.Admins
		if err := json.Unmarshal([]byte(community.Admins), &admins); err != nil {
			common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
			return
		}
		var tadmins models.Admins
		for _, admin := range admins {
			if admin == user.Username {
				continue
			}
			tadmins = append(tadmins, admin)
		}
		if tmpl, err := json.Marshal(&tadmins); err != nil {
			common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
			return
		} else {
			community.Admins = string(tmpl)
		}
		db.Where("id = ?", community.Id).Model(&models.Community{}).Update("admins", community.Admins)
		common.SuccessRes(ctx, nil, "退出成功")
		return
	}
	if len(community.Members) != 0 {
		var members models.Members
		if err := json.Unmarshal([]byte(community.Members), &members); err != nil {
			common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
			return
		}
		var tmembers models.Members
		for _, member := range members {
			if member == user.Username {
				continue
			}
			tmembers = append(tmembers, member)
		}
		if tmpl, err := json.Marshal(&tmembers); err != nil {
			common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
			return
		} else {
			community.Members = string(tmpl)
		}
		db.Where("id = ?", community.Id).Model(&models.Community{}).Update("members", community.Members)
		common.SuccessRes(ctx, nil, "退出成功")
		return
	}
}
