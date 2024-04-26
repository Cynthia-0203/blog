package controllers

import (
    "fmt"
    "github.com/Cynthia/goblog/pkg/flash"
    "github.com/Cynthia/goblog/pkg/logger"
    "net/http"

    "gorm.io/gorm"
)


type BaseController struct {
}


func (bc BaseController) ResponseForSQLError(w http.ResponseWriter, err error) {
    if err == gorm.ErrRecordNotFound {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprint(w, "404 falied to find article...")
    } else {
        logger.LogError(err)
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, "500 server internal error...")
    }
}

func (bc BaseController) ResponseForUnauthorized(w http.ResponseWriter, r *http.Request) {
    flash.Warning("unauthorized operation!")
    http.Redirect(w, r, "/", http.StatusFound)
}