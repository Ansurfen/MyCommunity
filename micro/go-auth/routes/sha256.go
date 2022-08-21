package routes

import (
	"go-auth/common"
	"go-auth/utils"

	"github.com/gin-gonic/gin"
)

func EncodeWithSHA256(ctx *gin.Context) {
	common.SuccessRes(ctx, gin.H{"data": utils.SHA256(ctx.PostForm("data"))}, "")
}
