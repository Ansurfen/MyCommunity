package community

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Del(ctx *gin.Context) {
	name := ctx.PostForm("cname")
	var community models.Community
	db := sql.GetDB("general")
	db.Where("name = ?", name).First(&community)
	var admins models.CommunityAdmin
	if err := json.Unmarshal([]byte(community.Admin), &admins); err != nil {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "System error")
		return
	}
	var user models.User
	if tmpl, b := ctx.Get("user"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "User don't exist")
		return
	} else {
		user = tmpl.(models.User)
	}
	if user.Username != admins.Host {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "权限不足")
		return
	}
	db.Delete(&community)
	common.SuccessRes(ctx, nil, "del successfully")
}
