package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strings"
	"vvvstore/internal/pkg/errno"
	"vvvstore/internal/pkg/jwt"
	"vvvstore/internal/pkg/response"
)

func JwtAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Authorization := ctx.Request.Header.Get("Authorization")
		if Authorization == "" {
			response.Fail(ctx, errno.UnAuthorizedError)
			ctx.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(Authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Fail(ctx, errno.UnAuthorizedError)
			ctx.Abort()
			return
		}

		// 验证token
		claims, err  := jwt.ParseToken(parts[1], viper.GetString("jwtSecret"))
		if err != nil {
			response.Fail(ctx, errno.UnAuthorizedError)
			ctx.Abort()
			return
		}

		ctx.Set("ID", claims.ID)
		ctx.Next()
	}
}
