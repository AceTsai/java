package middleware

import (
	"github.com/gin-gonic/gin"
	"im/tool"
)

// JwtAuth JWT认证中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			tool.Failed("未登录或非法访问", c)
			c.Abort()
			return
		}
		if err := tool.VerifyToken(token); err != nil {
			tool.Failed("登录已过期，请重新登录", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
