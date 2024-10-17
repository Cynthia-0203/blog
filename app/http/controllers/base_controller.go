package controllers

import (
	"net/http"
	"github.com/Cynthia/goblog/pkg/flash"
	"github.com/Cynthia/goblog/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type BaseController struct {
}


func (bc BaseController) ResponseForSQLError(c *gin.Context, err error) {
    if err == gorm.ErrRecordNotFound {
        c.Writer.WriteHeader(http.StatusNotFound)
        c.String(http.StatusNotFound, "404 未找到文章...")
    } else {
        logger.LogError(err)
        c.Writer.WriteHeader(http.StatusInternalServerError)
        c.String(http.StatusInternalServerError, "500 服务器内部错误...")
    }
}

func (bc BaseController) ResponseForUnauthorized(c *gin.Context) {
    flash.Warning(c,"unauthorized operation!")
    c.Redirect(http.StatusFound, "/")
}