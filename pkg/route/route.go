package route

import (
	"fmt"
	"strings"

	"github.com/Cynthia/goblog/pkg/config"
)

var routeMap = map[string]string{
	"articles.index":    "/articles",
	"articles.show":     "/articles/:id",
	"articles.create":   "/articles/create",
	"articles.store":    "/articles",
	"articles.edit":     "/articles/:id/edit",
	"articles.delete":"/articles/:id/delete",

	"categories.show":    "/categories/:id",
	"categories.store":   "/categories",
	"categories.create": "/categories/create",

	"users.show":     "/user/:id",

	"auth.register":"/auth/register",
	"auth.login":"/auth/login",
	"auth.dologin":"/auth/dologin",
	"auth.doregister":"/auth/doregister",
	"auth.logout":"/auth/logout",
}

func Name2URL(routeName string, pairs ...string) string {
	fmt.Println(routeName,pairs)
	path, exists := routeMap[routeName]
	if !exists {
		panic("路由名称不存在，请检查 routes/web.go 文件")
	}

	// 替换路径中的参数
	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]
		value := pairs[i+1]
		path = replacePathParam(path, key, value)
	}
	fmt.Printf("path:%v",path)
	return config.GetString("app.url") + path
}
func replacePathParam(path, key, value string) string {
	return strings.Replace(path, ":"+key, value, -1)
}