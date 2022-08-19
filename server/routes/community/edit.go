package community

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Edit(ctx *gin.Context) {
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
	/*
		context: 保存着要更改的所有数据
	*/
	var post models.Post
	if tmpl, b := ctx.Get("post"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Post don't exist")
		return
	} else {
		post = tmpl.(models.Post)
	}
	var data models.Community
	if err := json.Unmarshal([]byte(post.Context), &data); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	if data.Name != "" {
		community.Name = data.Name
	}
	if data.Tags != "" {
		community.Tags = data.Tags
	}
	if data.Context != "" {
		community.Context = data.Context
	}
	db := sql.GetDB("general")
	db.Where("id = ?", community.Id).Save(&data)
	common.SuccessRes(ctx, nil, "更改成功")
}
