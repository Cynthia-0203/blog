package routes

import (
	// "github.com/Cynthia/goblog/app/http/middlewares"
	"github.com/Cynthia/goblog/app/http/controllers"
	"github.com/Cynthia/goblog/app/http/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *gin.Engine) {
	// r.Use(middlewares.StartSession)
	r.Static("/css", "./public/css")
	r.Static("/js", "./public/js")
    // 静态页面
    pc := new(controllers.PagesController)
    r.NoRoute(pc.NotFound)
    r.GET("/", pc.Home)


	ac := new(controllers.ArticlesController)
    // 不需要认证的路由
    r.GET("/articles", ac.Index)              // 显示所有文章
    r.GET("/articles/:id", ac.Show)           // 显示单篇文章

    // 需要认证的路由
    r.GET("/articles/create", middlewares.Auth(), ac.Create)  // 创建文章页面
    r.POST("/articles", middlewares.Auth(), ac.Store)         // 创建文章
    r.GET("/articles/:id/edit", middlewares.Auth(), ac.Edit)  // 编辑文章页面
    r.POST("/articles/:id", middlewares.Auth(), ac.Update)    // 更新文章
    r.POST("/articles/:id/delete", middlewares.Auth(), ac.Delete)  // 删除文章


	 // 用户认证
	auc := new(controllers.AuthController)
    r.GET("/auth/register", middlewares.Guest(), auc.Register)
	r.POST("/auth/doregister", middlewares.Guest(), auc.DoRegister)
	r.GET("/auth/login", middlewares.Guest(), auc.Login)
	r.POST("/auth/dologin", middlewares.Guest(), auc.DoLogin)
	r.POST("/auth/logout", middlewares.Auth(), auc.Logout)

	
    uc := new(controllers.UserController)
    r.GET("/user/:id", uc.Show)

	// 文章分类
	cc := new(controllers.CategoriesController)
	r.GET("/categories/create", middlewares.Auth(), cc.Create)
	r.POST("/categories", middlewares.Auth(), cc.Store)
	r.GET("/categories/:id", cc.Show)
}

