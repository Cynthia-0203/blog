package controllers

import (
	"fmt"
	"net/http"

	"github.com/Cynthia/goblog/app/models/article"
	"github.com/Cynthia/goblog/app/models/category"
	"github.com/Cynthia/goblog/app/requests"
	"github.com/Cynthia/goblog/pkg/flash"
	"github.com/Cynthia/goblog/pkg/route"
	"github.com/Cynthia/goblog/pkg/view"
	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
    BaseController
}


func (*CategoriesController) Create(c *gin.Context) {
    view.Render(c, view.D{}, "categories.create")
}


func (*CategoriesController) Store(c *gin.Context) {

    _category := category.Category{
        Name: c.PostForm("name"),
    }

   
    errors := requests.ValidateCategoryForm(_category)

    if len(errors) == 0 {
       
        _category.Create()
        if _category.ID > 0 {
			flash.Success(c,"分类创建成功")
            index:=route.Name2URL("categories.show", "id", _category.GetStringID())
            fmt.Println(index)
            c.Redirect(http.StatusFound, index)
        } else {
            c.String(http.StatusInternalServerError, "创建文章分类失败，请联系管理员")
        }
    } else {
        view.Render(c, view.D{
            "Category": _category,
            "Errors":   errors,
        }, "categories.create")
    }
}


func (cc *CategoriesController) Show(c *gin.Context) {
      
    id := c.Param("id")
    fmt.Println(id)
    
    _category, err := category.Get(id)
    if err!=nil{
        return
    }
    
    articles, pagerData, err := article.GetByCategoryID(_category.GetStringID(), c, 2)
  
    if err != nil {
        cc.ResponseForSQLError(c, err)
    } else {
        view.Render(c, view.D{
            "Articles":  articles,
            "PagerData": pagerData,
        }, "articles.index", "articles._article_meta")
    }
}