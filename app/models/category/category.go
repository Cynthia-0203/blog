package category

import (
	"github.com/Cynthia/goblog/app/models"
	"github.com/Cynthia/goblog/pkg/route"
)

// Category 文章分类
type Category struct {
    models.BaseModel

    Name string `gorm:"type:varchar(255);not null;" valid:"name"`
}

func (category Category) Link() string {
    return route.Name2URL("categories.show", "id", category.GetStringID())
}