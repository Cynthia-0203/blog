package bootstrap

import (
	"github.com/Cynthia/goblog/pkg/session"
	"github.com/Cynthia/goblog/routes"
	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
    router := gin.Default()
    router.Use(session.SetupSessionMiddleware("gblog"))
    routes.RegisterWebRoutes(router)
    return router
}


