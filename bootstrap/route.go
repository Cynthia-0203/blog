package bootstrap

import (
	"github.com/Cynthia/goblog/pkg/route"
	"github.com/Cynthia/goblog/routes"
	"github.com/gorilla/mux"
)


func SetupRoute() *mux.Router {
    router := mux.NewRouter()
    routes.RegisterWebRoutes(router)

    route.SetRoute(router)

    return router
}
