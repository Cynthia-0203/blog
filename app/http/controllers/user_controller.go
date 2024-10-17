package controllers

import (
	"net/http"
	"github.com/Cynthia/goblog/app/models/article"
	"github.com/Cynthia/goblog/app/models/user"
	"github.com/Cynthia/goblog/pkg/logger"
	"github.com/Cynthia/goblog/pkg/view"
	"github.com/gin-gonic/gin"
)


type UserController struct {
    BaseController
}


func (uc *UserController) Show(c *gin.Context) {

    id := c.Param("id")
    _user, err := user.Get(id)

    
    if err != nil {
        uc.ResponseForSQLError(c, err)
    } else {
        articles, err := article.GetByUserID(_user.GetStringID())
        if err != nil {
            logger.LogError(err)
            c.String(http.StatusInternalServerError, "500 server internal error...")
        } else {
            view.Render(c, view.D{
                "Articles": articles,
            }, "articles.index", "articles._article_meta")
        }
    }
}