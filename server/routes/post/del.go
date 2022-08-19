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

/*
	username from jwt
	from -> id
	to -> timestamp
	context
	timestamp -> level
*/
// API
func DelSubComment(ctx *gin.Context) {
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
	var tsubs []models.SubComment
	for _, v := range subs {
		if v.Timestamp == sub.To && v.From == user.Username {
			continue
		}
		tsubs = append(tsubs, v)
	}
	if tmpl, err := json.Marshal(&tsubs); err != nil {
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
	common.SuccessRes(ctx, nil, "删除成功")
}

func DelComment(ctx *gin.Context) {
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
	comments = append(comments[:level], comments[level+1:]...)
	if tmpl, err := json.Marshal(&comments); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	} else {
		post.Comments = string(tmpl)
	}
	db.Where("id = ?", id).Model(&models.Posts{}).Update("comments", post.Comments)
	common.SuccessRes(ctx, nil, "删除成功")
}
