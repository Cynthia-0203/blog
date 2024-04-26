package controllers

import (
    "fmt"
    "github.com/Cynthia/goblog/app/models/article"
    "github.com/Cynthia/goblog/app/models/user"
    "github.com/Cynthia/goblog/pkg/logger"
    "github.com/Cynthia/goblog/pkg/route"
    "github.com/Cynthia/goblog/pkg/view"
    "net/http"
)


type UserController struct {
    BaseController
}


func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {


    id := route.GetRouteVariable("id", r)
    _user, err := user.Get(id)

    
    if err != nil {
        uc.ResponseForSQLError(w, err)
    } else {
         
        articles, err := article.GetByUserID(_user.GetStringID())
        if err != nil {
            logger.LogError(err)
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprint(w, "500 server internal error...")
        } else {
            view.Render(w, view.D{
                "Articles": articles,
            }, "articles.index", "articles._article_meta")
        }
    }
}