package controllers

import (
	"fmt"
	"net/http"

	"github.com/Cynthia/goblog/app/models/article"
	"github.com/Cynthia/goblog/app/policies"
	"github.com/Cynthia/goblog/app/requests"
	"github.com/Cynthia/goblog/pkg/auth"
	"github.com/Cynthia/goblog/pkg/route"
	"github.com/Cynthia/goblog/pkg/view"
)


type ArticlesController struct {
    BaseController
}

func (ac *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {

 
    id := route.GetRouteVariable("id", r)
    article, err := article.Get(id)
    if err != nil {
        ac.ResponseForSQLError(w, err)
    } else {
        
        view.Render(w, view.D{
            "Article":          article,
            "CanModifyArticle": policies.CanModifyArticle(article),
        }, "articles.show", "articles._article_meta")
    }
}

func (ac *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {

    
    articles, pagerData,err := article.GetAll(r,2)

    if err != nil {
        ac.ResponseForSQLError(w, err)
    } else {

       
        view.Render(w, view.D{
            "Articles":  articles,
            "PagerData": pagerData,
        }, "articles.index", "articles._article_meta")
    }
}


func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
    view.Render(w, view.D{}, "articles.create", "articles._form_field")
}


func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
    
    currentUser := auth.User()
    _article := article.Article{
        Title:  r.PostFormValue("title"),
        Body:   r.PostFormValue("body"),
        UserID: currentUser.ID,
    }

    errors := requests.ValidateArticleForm(_article)

    if len(errors) == 0 {
        _article.Create()
        if _article.ID > 0 {
            indexURL := route.Name2URL("articles.show", "id", _article.GetStringID())
            http.Redirect(w, r, indexURL, http.StatusFound)
        } else {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprint(w, "failed to create article...")
        }
    } else {
        view.Render(w, view.D{
            "Article": _article,
            "Errors":  errors,
        }, "articles.create", "articles._form_field")
    }
}


func (ac *ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {

    
    id := route.GetRouteVariable("id", r)
    _article, err := article.Get(id)
    if err != nil {
        ac.ResponseForSQLError(w, err)
    } else {

        if !policies.CanModifyArticle(_article) {
            ac.ResponseForUnauthorized(w, r)
        } else {
            view.Render(w, view.D{
                "Article": _article,
                "Errors":  view.D{},
            }, "articles.edit", "articles._form_field")
        }
    }
}


func (ac *ArticlesController) Update(w http.ResponseWriter, r *http.Request) {

    id := route.GetRouteVariable("id", r)
    _article, err := article.Get(id)
    if err != nil {
        ac.ResponseForSQLError(w, err)
    } else {
        if !policies.CanModifyArticle(_article) {
            ac.ResponseForUnauthorized(w, r)
        } else {

            _article.Title = r.PostFormValue("title")
            _article.Body = r.PostFormValue("body")

            errors := requests.ValidateArticleForm(_article)

            if len(errors) == 0 {

                rowsAffected, err := _article.Update()
                if err != nil {
                    
                    w.WriteHeader(http.StatusInternalServerError)
                    fmt.Fprint(w, "500 sever internal error")
                    return
                }

                if rowsAffected > 0 {
                    showURL := route.Name2URL("articles.show", "id", id)
                    http.Redirect(w, r, showURL, http.StatusFound)
                } else {
                    fmt.Fprint(w, "no change!")
                }
            } else {
                view.Render(w, view.D{
                    "Article": _article,
                    "Errors":  errors,
                }, "articles.edit", "articles._form_field")
            }
        }
    }
}

func (ac *ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {

    id := route.GetRouteVariable("id", r)
    _article, err := article.Get(id)
    if err != nil {
        ac.ResponseForSQLError(w, err)
    } else {
        if !policies.CanModifyArticle(_article) {
            ac.ResponseForUnauthorized(w, r)
        } else {
            
            rowsAffected, err := _article.Delete()
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                fmt.Fprint(w, "500 server internal error...")
            } else {
               
                if rowsAffected > 0 {
                    indexURL := route.Name2URL("articles.index")
                    http.Redirect(w, r, indexURL, http.StatusFound)
                } else {
    
                    w.WriteHeader(http.StatusNotFound)
                    fmt.Fprint(w, "404 not find article...")
                }
            }
        }
    }
}