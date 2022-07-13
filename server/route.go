package server

import (
	"MyCommunity/middlewares"
	"MyCommunity/models"
	"MyCommunity/routes/admin"
	"MyCommunity/routes/commit"
	"MyCommunity/routes/community"
	"MyCommunity/routes/user"
)

func (s *SEngine) UseRouter() *SEngine {
	s.POST("/user/login", user.Login)
	s.POST("/user/register", middlewares.CheckUniqueUsername(), middlewares.CheckUniqueKey(), user.Register)
	s.POST("/user/edit/info", user.EditInfo)
	s.POST("/user/edit/pwd", user.EditPwd)
	s.POST("/user/find", user.Find)
	s.DELETE("/user/del", user.Cancel)
	s.GET("/user/info", user.Info)

	s.POST("/community/new", middlewares.AuthJWT(), community.New)
	// 验证权限
	s.POST("/community/edit", middlewares.AuthJWT(), middlewares.CheckCommunityRight(), community.Edit)
	// 验证权限
	s.POST("/community/del", middlewares.AuthJWT(), community.Del)
	s.GET("/community/info", community.Info)

	commitRouter := s.Group("/commit", middlewares.CheckRight(models.Admin))
	{
		commitRouter.GET("/info", commit.Info)
		commitRouter.POST("/pass", commit.Pass)
	}

	// 后台
	s.GET("/admin/info", middlewares.CheckRight(models.ROOT), admin.Info)
	s.POST("/admin/add", middlewares.CheckRight(models.ROOT), admin.Add)
	s.POST("/admin/del", middlewares.CheckRight(models.ROOT), admin.Del)
	s.POST("/admin/find", middlewares.CheckRight(models.ROOT), admin.Find)
	return s
}
