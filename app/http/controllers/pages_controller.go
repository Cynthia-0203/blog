package controllers

import (
    "fmt"
    "net/http"
)


type PagesController struct {
}


func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "<h1>Hello, welcome</h1>")
}


func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
        "<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}


func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprint(w, "<h1>Don't find page!</p>")
}