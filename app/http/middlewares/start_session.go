package middlewares

import (
	"net/http"

	"github.com/Cynthia/goblog/pkg/session"
)


func StartSession(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

       
        session.StartSession(w, r)
        next.ServeHTTP(w, r)
    })
}