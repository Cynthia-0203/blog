package article

import (
	"github.com/Cynthia/goblog/pkg/logger"
	"github.com/Cynthia/goblog/pkg/model"
	"github.com/Cynthia/goblog/pkg/pagination"
	"github.com/Cynthia/goblog/pkg/route"
	"github.com/gin-gonic/gin"

	"github.com/Cynthia/goblog/pkg/types"
)

// Get 通过 ID 获取文章
func Get(idstr string) (Article, error) {
    var article Article
    id := types.StringToUint64(idstr)
    if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
        return article, err
    }

    return article, nil
}


func GetAll(c *gin.Context, perPage int) ([]Article, pagination.ViewData, error) {

    
    db := model.DB.Model(Article{}).Order("created_at desc")
    _pager := pagination.New(c.Request, db, route.Name2URL("home"), perPage)

   
    viewData := _pager.Paging()


    var articles []Article
    _pager.Results(&articles)

    return articles, viewData, nil
}


func (article *Article) Create() (err error) {
    if err = model.DB.Create(&article).Error; err != nil {
        logger.LogError(err)
        return err
    }

    return nil
}
func (article *Article) Update() (rowsAffected int64, err error) {
    result := model.DB.Save(&article)
    if err = result.Error; err != nil {
        logger.LogError(err)
        return 0, err
    }

    return result.RowsAffected, nil
}


func (article *Article) Delete() (rowsAffected int64, err error) {
    result := model.DB.Delete(&article)
    if err = result.Error; err != nil {
        logger.LogError(err)
        return 0, err
    }

    return result.RowsAffected, nil
}

func GetByUserID(uid string) ([]Article, error) {
    var articles []Article
    if err := model.DB.Where("user_id = ?", uid).Preload("User").Find(&articles).Error; err != nil {
        return articles, err
    }
    return articles, nil
}

func GetByCategoryID(cid string, c *gin.Context, perPage int) ([]Article, pagination.ViewData, error) {

    
    db := model.DB.Model(Article{}).Where("category_id = ?", cid).Order("created_at desc")
    _pager := pagination.New(c.Request, db, route.Name2URL("categories.show", "id", cid), perPage)

    
    viewData := _pager.Paging()

    
    var articles []Article
    _pager.Results(&articles)

    return articles, viewData, nil
}