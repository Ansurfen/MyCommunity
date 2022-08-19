package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	DEBUG = iota
	STANDARD
)

var MODE int

func init() {
	MODE = DEBUG
}

func SetResponseMode(mode int) {
	MODE = mode
}

func responseHelper(ctx *gin.Context, httpStatus int, data gin.H, msg string, debug bool) {
	ctx.JSON(httpStatus, gin.H{
		"data": data,
		"msg":  msg,
	})
	if debug {
		logrus.Info("[DEBUG]" + msg)
	}
}

func stdResponse(ctx *gin.Context, httpStatus int, data gin.H, msg string) {
	responseHelper(ctx, httpStatus, data, msg, false)
}

func debugResponse(ctx *gin.Context, httpStatus int, data gin.H, msg string) {
	responseHelper(ctx, httpStatus, data, msg, true)
}

func CommonRes(ctx *gin.Context, httpStatus int, data gin.H, msg string) {
	switch MODE {
	case DEBUG:
		debugResponse(ctx, httpStatus, data, msg)
	case STANDARD:
		stdResponse(ctx, httpStatus, data, msg)
	default:
		logrus.Fatal("Fail to format response.")
	}
}

func SuccessRes(ctx *gin.Context, data gin.H, msg string) {
	switch MODE {
	case DEBUG:
		debugResponse(ctx, http.StatusOK, data, msg)
	case STANDARD:
		stdResponse(ctx, http.StatusOK, data, msg)
	default:
		logrus.Fatal("Fail to format response.")
	}
}

func FailRes(ctx *gin.Context, data gin.H, msg string) {
	switch MODE {
	case DEBUG:
		debugResponse(ctx, http.StatusBadRequest, data, msg)
	case STANDARD:
		stdResponse(ctx, http.StatusBadRequest, data, msg)
	default:
		logrus.Fatal("Fail to format response.")
	}
}
