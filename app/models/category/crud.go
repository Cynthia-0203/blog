package category

import (
	"github.com/Cynthia/goblog/pkg/logger"
	"github.com/Cynthia/goblog/pkg/model"

	// "github.com/Cynthia/goblog/pkg/route"
	"github.com/Cynthia/goblog/pkg/types"
)


func (category *Category) Create() (err error) {
    if err = model.DB.Create(&category).Error; err != nil {
        logger.LogError(err)
        return err
    }
    return nil
}

func All() ([]Category, error) {
    var categories []Category
    if err := model.DB.Find(&categories).Error; err != nil {
        return categories, err
    }
    return categories, nil
}


func Get(idstr string) (Category, error) {
    var category Category
    id := types.StringToUint64(idstr)
    if err := model.DB.First(&category, id).Error; err != nil {
        return category, err
    }

    return category, nil
}