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

func AddPost(ctx *gin.Context) {
	var post models.Post
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	if tmpl, b := ctx.Get("post"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Post don't exist")
		return
	} else {
		post = tmpl.(models.Post)
		post.Score = 0
		post.Author = user.Username
		post.Timestamp = utils.NowTimestampByString()
		post.Id = utils.RandValue(post.Id, post.Timestamp, post.Author)
	}
	var community models.Community
	if tmpl, b := ctx.Get("community"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community don't exist")
		return
	} else {
		community = tmpl.(models.Community)
	}
	if len(community.Posts) == 0 {
		community.Posts = "[]"
	}
	var posts []models.Post
	if err := json.Unmarshal([]byte(community.Posts), &posts); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	posts = append(posts, post)
	if tmpl, err := json.Marshal(&posts); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	} else {
		community.Posts = string(tmpl)
	}
	db := sql.GetDB("general")
	db.Where("id = ?", community.Id).Model(&models.Community{}).Update("posts", community.Posts)
	ps := models.NewPosts(post.Id, post.Title, post.Author, post.Timestamp, post.Context, post.Tags)
	db.Insert(&ps)
	common.SuccessRes(ctx, nil, "发帖成功")
}

func DelPost(ctx *gin.Context) {

}

func EditPost(ctx *gin.Context) {

}

func InfoPost(ctx *gin.Context) {
	var community models.Community
	db := sql.GetDB("general")
	db.Where("name = ?", ctx.PostForm("cname")).First(&community)
	if community.Id == "" {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community don't exist")
		return
	}
	common.SuccessRes(ctx, gin.H{"data": community.Posts}, "获取帖子成功")
}
