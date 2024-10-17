package middlewares

import (
	"net/http"

	"github.com/Cynthia/goblog/pkg/auth"
	"github.com/Cynthia/goblog/pkg/flash"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        if !auth.Check(c) {
            // 设置提示消息 (根据你在 Gin 中如何处理闪存消息来进行调整)
            flash.Warning(c,"must login!")
            // 如果身份验证失败，则重定向到首页
            c.Redirect(http.StatusFound, "/")
            // 终止请求，不再执行后续的处理器
            c.Abort()
            return
        }

        // 如果已通过身份验证，继续执行下一个处理器
        c.Next()
    }
}
