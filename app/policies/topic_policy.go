package policies

import (
    "github.com/Cynthia/goblog/app/models/article"
    "github.com/Cynthia/goblog/pkg/auth"
)


func CanModifyArticle(_article article.Article) bool {
    return auth.User().ID == _article.UserID
}