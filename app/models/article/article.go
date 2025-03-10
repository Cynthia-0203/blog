package article

import (
	"strconv"

	"github.com/Cynthia/goblog/app/models"
	"github.com/Cynthia/goblog/app/models/user"
	"github.com/Cynthia/goblog/pkg/route"
)


type Article struct {
    models.BaseModel
	UserID uint64 `gorm:"not null;index"`
    User   user.User
    Title string `gorm:"type:varchar(255);not null;" valid:"title"`
    Body  string `gorm:"type:longtext;not null;" valid:"body"`
	CategoryID uint64 `gorm:"not null;index"`
}


// Link 方法用来生成文章链接
func (article Article) Link() string {
    return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}

func (article Article) CreatedAtDate() string {
    return article.CreatedAt.Format("2006-01-02")
}