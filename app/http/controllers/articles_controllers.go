package controllers

import (
	"net/http"

	"github.com/Cynthia/goblog/app/models/article"
	"github.com/Cynthia/goblog/app/policies"
	"github.com/Cynthia/goblog/app/requests"
	"github.com/Cynthia/goblog/pkg/auth"
	"github.com/Cynthia/goblog/pkg/route"
	"github.com/Cynthia/goblog/pkg/view"
	"github.com/gin-gonic/gin"
)


type ArticlesController struct {
    BaseController
}

func (ac *ArticlesController) Show(c *gin.Context) {
    id := c.Param("id")
    article, err := article.Get(id)
    if err != nil {
        ac.ResponseForSQLError(c, err)
    } else {
        
        view.Render(c, view.D{
            "Article":          article,
            "CanModifyArticle": policies.CanModifyArticle(c,article),
        }, "articles.show", "articles._article_meta")
    }
}

func (ac *ArticlesController) Index(c *gin.Context) {

    
    articles, pagerData,err := article.GetAll(c,2)

    if err != nil {
        ac.ResponseForSQLError(c, err)
    } else {
        view.Render(c, view.D{
            "Articles":  articles,
            "PagerData": pagerData,
        }, "articles.index", "articles._article_meta")
    }
}


func (*ArticlesController) Create(c *gin.Context) {
    view.Render(c, view.D{}, "articles.create", "articles._form_field")
}


func (*ArticlesController) Store(c *gin.Context) {
    
    currentUser := auth.User(c)
    _article := article.Article{
        Title:  c.PostForm("title"),
        Body:   c.PostForm("body"),
        UserID: currentUser.ID,
    }

    errors := requests.ValidateArticleForm(_article)

    if len(errors) == 0 {
        _article.Create()
        if _article.ID > 0 {
            index:=route.Name2URL("articles.show", "id", _article.GetStringID())
            c.Redirect(http.StatusFound, index)
        } else {
            c.String(http.StatusInternalServerError, "failed to create article...")
        }
    } else {
        view.Render(c, view.D{
            "Article": _article,
            "Errors":  errors,
        }, "articles.create", "articles._form_field")
    }
}


func (ac *ArticlesController) Edit(c *gin.Context) {
    id := c.Param("id")
    _article, err := article.Get(id)
    if err != nil {
        ac.ResponseForSQLError(c, err)
    } else {

        if !policies.CanModifyArticle(c,_article) {
            ac.ResponseForUnauthorized(c)
        } else {
            view.Render(c, view.D{
                "Article": _article,
                "Errors":  view.D{},
            }, "articles.edit", "articles._form_field")
        }
    }
}


func (ac *ArticlesController) Update(c *gin.Context) {

    id := c.Param("id")
    _article, err := article.Get(id)
    if err != nil {
        ac.ResponseForSQLError(c, err)
    } 
    if !policies.CanModifyArticle(c,_article) {
        ac.ResponseForUnauthorized(c)
    } else {

        _article.Title = c.PostForm("title")
        _article.Body = c.PostForm("body")


        errors := requests.ValidateArticleForm(_article)

        if len(errors) == 0 {

            rowsAffected, err := _article.Update()
            if err != nil {
                c.String(http.StatusInternalServerError, "500 server internal error")
                return
            }

            if rowsAffected > 0 {
                index:=route.Name2URL("articles.show", "id", _article.GetStringID())
                c.Redirect(http.StatusFound, index)
            } else {
                c.String(http.StatusOK, "no change!")
            }
        } else {
            view.Render(c, view.D{
                "Article": _article,
                "Errors":  errors,
            }, "articles.edit", "articles._form_field")
        }
        
    }
}

func (ac *ArticlesController) Delete(c *gin.Context) {

    id := c.Param("id")
    _article, err := article.Get(id)
    if err != nil {
        ac.ResponseForSQLError(c, err)
    } else {
        if !policies.CanModifyArticle(c,_article) {
            ac.ResponseForUnauthorized(c)
        } else {
            
            rowsAffected, err := _article.Delete()
            if err != nil {
                c.String(http.StatusInternalServerError, "500 server internal error...")
            } else {
                if rowsAffected > 0 {
                    c.Redirect(http.StatusFound, "/articles/index")
                } else {
                    c.String(http.StatusNotFound, "404 not find article...")
                }
            }
        }
    }
}