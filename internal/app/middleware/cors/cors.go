package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"vvvstore/internal/pkg/logger"
)

func New() gin.HandlerFunc {
	if viper.GetBool("app.cors") {
		return cors.Default()
	} else {
		return func(ctx *gin.Context) {
			logger.Debug("cros is not enabled")
			ctx.Next()
		}
	}
}
