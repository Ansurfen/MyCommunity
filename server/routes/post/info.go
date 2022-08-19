package post

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context) {
	var post models.Posts
	db := sql.GetDB("general")
	db.Where("id = ?", ctx.PostForm("id")).First(&post)
	if post.Title == "" {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Post don't exist")
		return
	}
	res, err := json.Marshal(&post)
	if err != nil {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "系统错误")
		return
	}
	common.SuccessRes(ctx, gin.H{"data": string(res)}, "获取帖子成功")
}
