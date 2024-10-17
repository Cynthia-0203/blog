package policies

import (
	"github.com/Cynthia/goblog/app/models/article"
	"github.com/Cynthia/goblog/pkg/auth"
	"github.com/gin-gonic/gin"
)


func CanModifyArticle(c *gin.Context,_article article.Article) bool {
    return auth.User(c).ID == _article.UserID
}