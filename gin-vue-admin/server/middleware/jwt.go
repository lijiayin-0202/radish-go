package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/radish-go/gin-vue-admin/server/global"
	"github.com/radish-go/gin-vue-admin/server/model/common/response"
	"github.com/radish-go/gin-vue-admin/server/utils"
)

// JWTAuth JWT 鉴权中间件，校验通过后把 claims 写入上下文
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			response.FailWithCode(response.AUTHERROR, "请求未携带 token", c)
			c.Abort()
			return
		}
		token := strings.TrimPrefix(header, "Bearer ")

		claims, err := utils.ParseToken(token)
		if err != nil {
			global.GVA_LOG.Warn("token 校验失败", zap.Error(err))
			response.FailWithCode(response.AUTHERROR, "token 无效或已过期", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Set("userID", claims.ID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
