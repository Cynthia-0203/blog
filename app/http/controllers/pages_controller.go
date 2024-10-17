package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type PagesController struct {
}


func (*PagesController) Home(c *gin.Context) {
    // fmt.Fprint(c.Writer, "<h1>Hello, welcome</h1>")
    c.Redirect(http.StatusFound, "/articles")
}


func (*PagesController) NotFound(c *gin.Context) {
    c.Writer.WriteHeader(http.StatusNotFound)
    // 返回自定义的 HTML 内容
    c.Writer.Write([]byte("<h1>Page Not Found!</h1>"))  
}