package policies

import (
    "github.com/Cynthia/goblog/app/models/article"
    "github.com/Cynthia/goblog/pkg/auth"
)

// CanModifyArticle 是否允许修改话题
func CanModifyArticle(_article article.Article) bool {
    return auth.User().ID == _article.UserID
}