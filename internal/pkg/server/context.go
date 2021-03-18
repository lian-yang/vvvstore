package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"vvvstore/internal/pkg/response"
)

type Context struct {
	*gin.Context
}

type HandlerFunc func(ctx *Context)

// 获取日志记录器
func (ctx *Context) GetLogger() *zap.Logger {
	log, _ := ctx.Get("logger")
	return log.(*zap.Logger)
}

// 响应成功
func (ctx *Context) Success(msg string, data interface{}) {
	response.Success(ctx.Context, msg, data)
}

// 响应失败
func (ctx *Context) Fail(err error) {
	response.Fail(ctx.Context, err)
}

// gin.Context 包装
func (s *Server) Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			c,
		}
		h(ctx)
	}
}
