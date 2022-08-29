package post

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Edit(ctx *gin.Context) {
	var data models.Posts
	if err := ctx.ShouldBindJSON(&data); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}

	if user.Username != data.Author {
		common.FailRes(ctx, nil, "没有权限")
		return
	}
	db := sql.GetDB("general")
	db.Where("id = ?", data.Id).Model(&models.Posts{}).Update(gin.H{
		"context": data.Context,
		"tags":    data.Tags,
	})
	common.SuccessRes(ctx, nil, "修改成功")
}

/*
	from -> post id
	to ->
	context -> editted context
	timestamp -> level
*/
// API
func EditComment(ctx *gin.Context) {
	var sub models.SubComment
	if err := ctx.ShouldBindJSON(&sub); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	id := sub.From
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	var post models.Posts
	db := sql.GetDB("general")
	db.Where("id = ?", id).First(&post)
	if post.Title == "" {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Post don't exist")
		return
	}
	var comments []models.Comment
	if err := json.Unmarshal([]byte(post.Comments), &comments); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	level, _ := strconv.Atoi(sub.Timestamp)
	if len(comments) == 0 || len(comments)-1 < level {
		common.FailRes(ctx, nil, "System error")
		return
	}
	if comments[level].Host != user.Username {
		common.FailRes(ctx, nil, "没有权限")
		return
	}
	comments[level].Context = sub.Context
	if tmpl, err := json.Marshal(&comments); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	} else {
		post.Comments = string(tmpl)
	}
	db.Where("id = ?", id).Model(&models.Posts{}).Update("comments", post.Comments)
	common.SuccessRes(ctx, nil, "修改成功")
}

func EditSubComment(ctx *gin.Context) {
	var sub models.SubComment
	if err := ctx.ShouldBindJSON(&sub); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	id := sub.From
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	var post models.Posts
	db := sql.GetDB("general")
	db.Where("id = ?", id).First(&post)
	if post.Title == "" {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Post don't exist")
		return
	}
	var comments []models.Comment
	if err := json.Unmarshal([]byte(post.Comments), &comments); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	level, _ := strconv.Atoi(sub.Timestamp)
	if len(comments) == 0 || len(comments)-1 < level {
		common.FailRes(ctx, nil, "System error")
		return
	}
	var subs []models.SubComment
	if err := json.Unmarshal([]byte(comments[level].Sub), &subs); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	for k, v := range subs {
		if v.Timestamp == sub.To && v.From == user.Username {
			subs[k].Context = sub.Context
		}
	}
	if tmpl, err := json.Marshal(&subs); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	} else {
		comments[level].Sub = string(tmpl)
	}
	if tmpl, err := json.Marshal(&comments); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	} else {
		post.Comments = string(tmpl)
	}
	db.Where("id = ?", id).Model(&models.Posts{}).Update("comments", post.Comments)
	common.SuccessRes(ctx, nil, "修改成功")
}

// 得实现一个序列化
// 同时还得有个反序列化
