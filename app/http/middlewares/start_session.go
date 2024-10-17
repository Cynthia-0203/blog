package middlewares

import (
	"github.com/Cynthia/goblog/pkg/session"
	"github.com/gin-gonic/gin"
)


func StartSession() gin.HandlerFunc {
    return session.SetupSessionMiddleware("gblog")
}