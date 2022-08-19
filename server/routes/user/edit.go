package user

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditInfo(ctx *gin.Context) {
	var data models.User
	if err := ctx.ShouldBindJSON(&data); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	if data.Alias != "" {
		user.Alias = data.Alias
	}
	if data.Email != "" && common.CheckEmail(data.Email) {
		user.Email = utils.EncodeAESWithKey(utils.BackedKey, data.Email)
	}
	if data.Telephone != "" && common.CheckPhone(data.Telephone) {
		user.Telephone = utils.EncodeAESWithKey(utils.BackedKey, data.Telephone)
	}
	if data.School != "" {
		user.School = data.School
	}
	db := sql.GetDB("general")
	db.Update(&user)
	common.SuccessRes(ctx, nil, "Edit successfully")
}

func EditPwd(ctx *gin.Context) {
	oldPwd, newPwd := ctx.PostForm("oldpwd"), ctx.PostForm("newpwd")
	oldPwd, newPwd = utils.DecodeAESWithKey(utils.FrontKey, oldPwd), utils.DecodeAESWithKey(utils.FrontKey, newPwd)
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	if utils.EqualHashPassword(oldPwd, user.Password) != nil {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "old password err")
		return
	}
	user.Password = utils.HashPassword(newPwd)
	db := sql.GetDB("general")
	db.Update(&user)
	common.SuccessRes(ctx, nil, "Edit successfully")
}

func EditTelephone(ctx *gin.Context) {
	oldPhone, newPhone := ctx.PostForm("oldphone"), ctx.PostForm("newphone")
	oldPhone, newPhone = utils.DecodeAESWithKey(utils.FrontKey, oldPhone), utils.DecodeAESWithKey(utils.FrontKey, newPhone)
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	if utils.DecodeAESWithKey(utils.BackedKey, user.Telephone) != oldPhone {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Telephone error")
		return
	}
	user.Telephone = utils.EncodeAESWithKey(utils.BackedKey, newPhone)
	db := sql.GetDB("general")
	db.Update(&user)
	common.SuccessRes(ctx, nil, "Edit successfully")
}

func EditEmail(ctx *gin.Context) {
	oldEmail, newEmail := ctx.PostForm("oldemail"), ctx.PostForm("newemail")
	oldEmail, newEmail = utils.DecodeAESWithKey(utils.FrontKey, oldEmail), utils.DecodeAESWithKey(utils.FrontKey, newEmail)
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	if utils.DecodeAESWithKey(utils.BackedKey, user.Email) != oldEmail {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Email error")
		return
	}
	user.Email = utils.EncodeAESWithKey(utils.BackedKey, newEmail)
	db := sql.GetDB("general")
	db.Update(&user)
	common.SuccessRes(ctx, nil, "Edit successfully")
}
