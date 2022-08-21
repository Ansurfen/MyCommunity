package routes

import (
	"go-auth/common"
	"go-auth/utils"

	"github.com/gin-gonic/gin"
)

func EncodeWithMD5(ctx *gin.Context) {
	println(utils.MD5(ctx.PostForm("data")))
	common.SuccessRes(ctx, gin.H{"data": utils.MD5(ctx.PostForm("data"))}, "")
}
