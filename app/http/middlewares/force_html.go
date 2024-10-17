package middlewares

import (
	"github.com/gin-gonic/gin"
)


func ForceHTML() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 设置 Content-Type 为 text/html
        c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
        c.Next()
    }
}
