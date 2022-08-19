package server

import (
	"MyCommunity/utils"

	"github.com/gin-gonic/gin"
)

type SEngine struct {
	*gin.Engine
}

func NewApp() *SEngine {
	return &SEngine{
		Engine: gin.New(),
	}
}

func (s *SEngine) Default() *SEngine {
	s.Engine = gin.Default()
	return s
}

func (s *SEngine) Run() {
	s.Engine.Run(":" + utils.Conf.GetString("server.port"))
}