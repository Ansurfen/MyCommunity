package post

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	var comment models.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	var post models.Posts
	db := sql.GetDB("general")
	db.Where("id = ?", comment.Host).First(&post)
	if post.Title == "" {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Post don't exist")
		return
	}
	var user models.User
	if tmpl, b := ctx.Get("user"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "User don't exist")
		return
	} else {
		user = tmpl.(models.User)
	}
	comment.Host = user.Username
	comment.Timestamp = utils.NowTimestampByString()
	comment.Sub = "[]"
	if len(post.Comments) == 0 {
		post.Comments = "[]"
	}
	var comments []models.Comment
	if err := json.Unmarshal([]byte(post.Comments), &comments); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	comments = append(comments, comment)
	if tmpl, err := json.Marshal(&comments); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	} else {
		post.Comments = string(tmpl)
	}
	db.Where("id = ?", post.Id).Model(&models.Posts{}).Update("comments", post.Comments)
	common.SuccessRes(ctx, nil, "评论成功")
}
