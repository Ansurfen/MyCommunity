package commit

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pass(ctx *gin.Context) {
	cname := ctx.PostForm("cname")
	db := sql.GetDB("general")
	var community models.Community
	db.Where("name = ?", cname).First(&community)
	if community.Name == "" {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community don't exist")
		return
	}
	var admins models.CommunityAdmin
	if err := json.Unmarshal([]byte(community.Admin), &admins); err != nil {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "System error")
		return
	}
	members := models.NewCommunityMember(community.Name, admins, 100)
	db.Insert(members)
	community.Status = models.ACTIVE
	db.Update(&community)
	common.SuccessRes(ctx, nil, "审核通过")
}
