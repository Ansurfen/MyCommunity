package community

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddNote(ctx *gin.Context) {
	/*
		title -> title
		context -> context
	*/
	var post models.Post
	var note models.Notification
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	if tmpl, b := ctx.Get("post"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Post don't exist")
		return
	} else {
		post = tmpl.(models.Post)
		note.Title = post.Title
		note.Context = post.Context
		note.Timestamp = utils.NowTimestampByString()
	}
	var community models.Community
	if tmpl, b := ctx.Get("community"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community don't exist")
		return
	} else {
		community = tmpl.(models.Community)
	}
	if len(community.Notes) == 0 {
		community.Notes = "[]"
	}
	var notes []models.Notification
	if err := json.Unmarshal([]byte(community.Notes), &notes); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	notes = append(notes, note)
	if tmpl, err := json.Marshal(&notes); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	} else {
		community.Notes = string(tmpl)
	}
	db := sql.GetDB("general")
	db.Where("id = ?", community.Id).Model(&models.Community{}).Update("notes", community.Notes)
	common.SuccessRes(ctx, nil, "通知成功")
}

func InfoNote(ctx *gin.Context) {
	var community models.Community
	db := sql.GetDB("general")
	db.Where("name = ?", ctx.PostForm("cname")).First(&community)
	if community.Id == "" {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community don't exist")
		return
	}
	common.SuccessRes(ctx, gin.H{"data": community.Notes}, "获取通知成功")
}
