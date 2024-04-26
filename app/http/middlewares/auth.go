package middlewares

import (
    "github.com/Cynthia/goblog/pkg/auth"
    "github.com/Cynthia/goblog/pkg/flash"
    "net/http"
)

func Auth(next HttpHandlerFunc) HttpHandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        if !auth.Check() {
            flash.Warning("must login!")
            http.Redirect(w, r, "/", http.StatusFound)
            return
        }

        next(w, r)
    }
}