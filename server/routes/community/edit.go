package community

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

const savePath string = "./images/community/"

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
	if data.Tags != "" {
		community.Tags = data.Tags
	}
	if data.Context != "" {
		community.Context = data.Context
	}
	db := sql.GetDB("general")
	db.Model(&models.Community{}).Where("id = ?", community.Id).Updates(gin.H{
		"context": community.Context,
		"tags":    community.Tags,
	})
	common.SuccessRes(ctx, nil, "更改成功")
}

func Update(ctx *gin.Context) {
	if tmpl, b := ctx.Get("header"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Header don't exist")
		return
	} else {
		header := tmpl.(models.AuthHeader)
		db := sql.GetDB("general")
		var community models.Community
		db.Where("id = ?", header.First).First(&community)
		if community.Name == "" {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community don't exist")
			return
		}
		if file, err := ctx.FormFile("file"); err != nil {
			common.FailRes(ctx, nil, "上传失败")
		} else {
			ctx.SaveUploadedFile(file, savePath+community.Name+".png")
		}
		db.Where("id = ?", community.Id).Model(&models.Community{}).Update("image", savePath+community.Name+".png")
		common.SuccessRes(ctx, nil, "上传成功")
	}
}
