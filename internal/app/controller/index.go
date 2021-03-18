package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vvvstore/internal/pkg/server"
)

type IndexController struct {

}

// 构造函数
func NewIndexController() *IndexController {
	return &IndexController{}
}

// 注册路由
func (index *IndexController) Router (s *server.Server) {
	s.RouterGroup.GET( "/ping", index.Ping)
}

// Ping方法
// @Summary 实现Ping接口
// @Description 用于确认接口状态是否正常
// @Tags 系统相关接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} IndexController
// @Router /api/ping [get]
func (index *IndexController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}




