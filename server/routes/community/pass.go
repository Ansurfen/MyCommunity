package community

import (
	"github.com/gin-gonic/gin"
)

func Pass(ctx *gin.Context) {
	// cname := ctx.PostForm("cname")
	// wuser := ctx.PostForm("wuser")
	// var community models.Community
	// db := sql.GetDB("general")
	// db.Where("name = ?", cname).First(&community)
	// if community.Name != "" {
	// 	common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community don't exist")
	// 	return
	// }
	// var admins models.CommunityAdmin
	// if err := json.Unmarshal([]byte(community.Admin), &admins); err != nil {
	// 	common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "System error")
	// 	return
	// }
	// var adminslist []string
	// adminslist = append(adminslist, admins.Host)
	// adminslist = append(adminslist, admins.Admins...)
	// var user models.User
	// if tmpl, b := ctx.Get("user"); !b {
	// 	common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "User don't exist")
	// 	return
	// } else {
	// 	user = tmpl.(models.User)
	// }
	// hasRight := false
	// for _, admin := range adminslist {
	// 	if user.Username == admin {
	// 		hasRight = true
	// 	}
	// 	if wuser == admin {
	// 		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "用户已经在社团内")
	// 		return
	// 	}
	// }
	// if !hasRight {
	// 	common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "你没有权限操作")
	// 	return
	// }
	// var members models.CommunityMember
	// db.Where("name = ?", cname).First(&members)
	// for _, mem := range members.Members {
	// 	if mem == wuser {
	// 		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "用户已经在社团内")
	// 		return
	// 	}
	// }
	// members.Members = append(members.Members, wuser)
	// db.Update(&members)
	// common.SuccessRes(ctx, nil, "请求处理成功")
}
