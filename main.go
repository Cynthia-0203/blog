package main

import (
    "github.com/Cynthia/goblog/bootstrap"
    "github.com/Cynthia/goblog/config"
    c "github.com/Cynthia/goblog/pkg/config"
)

func init() {
    // 初始化配置信息
    config.Initialize()
}

func main() {
    // 初始化 SQL
    bootstrap.SetupDB()

    // 初始化路由绑定
    router := bootstrap.SetupRoute()

    router.Run(":" + c.GetString("app.port"))
}