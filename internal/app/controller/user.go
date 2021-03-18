package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"vvvstore/internal/pkg/errno"
	"vvvstore/internal/pkg/jwt"
	"vvvstore/internal/pkg/response"
	"vvvstore/internal/pkg/server"
)


type UserController struct {
	// 组合继承基础Controller
	*Controller
}

// 构造函数
func NewUserController() *UserController {
	return &UserController{}
}

// 注册路由
func (user *UserController) Router (s *server.Server) {
	s.RouterGroup.GET( "/ping", s.Handle(user.Ping))
	s.RouterGroup.GET( "/hello", s.Handle(user.Hello))
	s.GET("/api/v1/login", user.Login)
}

// 实现Ping业务
func (u *UserController) Ping(ctx *server.Context) {
	ctx.GetLogger().Info("ping")
	ctx.Success("ok", gin.H{
		"message": "pong",
	})
}

// 使用自定义的Context 使用server.Handle 包装 gin.Context
func (u *UserController) Hello(ctx *server.Context)  {
	ctx.Success("ok", gin.H{
		"messgae": u.hello(),
	})
}

// 登录
func (u *UserController) Login(ctx *gin.Context)  {
	secret :=  viper.GetString("jwtSecret")
	token, err := jwt.GenToken(1, secret)
	if err != nil {
		response.Fail(ctx, errno.GenTokenError)
		return
	}
	response.Success(ctx, "ok", gin.H{
		"token": token,
	})
}



