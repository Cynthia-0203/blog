package middlewares

import (
	"net/http"

	"github.com/Cynthia/goblog/pkg/auth"
	"github.com/Cynthia/goblog/pkg/flash"
	"github.com/gin-gonic/gin"
)


func Guest() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 检查用户是否已登录
        if auth.Check(c) {
            flash.Warning(c,"已登录用户无法访问此页面！")
            // 重定向到首页
            c.Redirect(http.StatusFound, "/")
            c.Abort()
            return
        }
        c.Next()
    }
}
