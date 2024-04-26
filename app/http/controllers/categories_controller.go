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
)


type CategoriesController struct {
    BaseController
}


func (*CategoriesController) Create(w http.ResponseWriter, r *http.Request) {
    view.Render(w, view.D{}, "categories.create")
}


func (*CategoriesController) Store(w http.ResponseWriter, r *http.Request) {

    _category := category.Category{
        Name: r.PostFormValue("name"),
    }

   
    errors := requests.ValidateCategoryForm(_category)

    if len(errors) == 0 {
       
        _category.Create()
        if _category.ID > 0 {
			flash.Success("分类创建成功")
            indexURL := route.Name2URL("categories.show", "id", _category.GetStringID())
            http.Redirect(w, r, indexURL, http.StatusFound)
        } else {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprint(w, "创建文章分类失败，请联系管理员")
        }
    } else {
        view.Render(w, view.D{
            "Category": _category,
            "Errors":   errors,
        }, "categories.create")
    }
}


func (cc *CategoriesController) Show(w http.ResponseWriter, r *http.Request) {
      
    id := route.GetRouteVariable("id", r)

      
    _category, err := category.Get(id)
    if err!=nil{
        return
    }
    
    articles, pagerData, err := article.GetByCategoryID(_category.GetStringID(), r, 2)
  
    if err != nil {
        cc.ResponseForSQLError(w, err)
    } else {
  
    
        view.Render(w, view.D{
            "Articles":  articles,
            "PagerData": pagerData,
        }, "articles.index", "articles._article_meta")
    }
}