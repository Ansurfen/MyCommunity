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
	var data models.Community
	if err := ctx.ShouldBindJSON(&data); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	db := sql.GetDB("general")
	var community models.Community
	db.Where("name = ?", data.Name).First(&community)
	if data.Context != "" {
		community.Context = data.Context
	}
	if data.Admin != "" {
		var admins models.CommunityAdmin
		if err := json.Unmarshal([]byte(community.Admin), &admins); err != nil {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "System error")
			return
		}
		community.Admin = data.Admin
	}
	common.SuccessRes(ctx, nil, "Edit successfully")
}
