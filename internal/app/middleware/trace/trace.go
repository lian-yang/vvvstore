package trace

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"vvvstore/internal/pkg/logger"
)

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.Request.Header.Get("X-Request-Id")

		if requestId == "" {
			u4 := uuid.NewV4()
			requestId = u4.String()
		}

		ctx.Set("X-Request-Id", requestId)
		ctx.Writer.Header().Set("X-Request-Id", requestId)

		logger := logger.Log.With(zap.String("x-request-id", requestId))
		ctx.Set("logger", logger)
		ctx.Next()
	}
}