package view

import (
	"html/template"
	"path/filepath"
	"strings"

	"github.com/Cynthia/goblog/app/models/category"
	"github.com/Cynthia/goblog/app/models/user"
	"github.com/Cynthia/goblog/pkg/auth"
	"github.com/Cynthia/goblog/pkg/flash"
	"github.com/Cynthia/goblog/pkg/logger"
	"github.com/Cynthia/goblog/pkg/route"

	"github.com/gin-gonic/gin"
)


type D map[string]interface{}


func Render(c *gin.Context, data D, tplFiles ...string) {
    RenderTemplate(c, "app", data, tplFiles...)
}

func RenderSimple(c *gin.Context, data D, tplFiles ...string) {
    RenderTemplate(c, "simple", data, tplFiles...)
}


func RenderTemplate(c *gin.Context, name string, data D, tplFiles ...string) {

    
    data["isLogined"] = auth.Check(c)
    data["flash"] = flash.All(c)
    data["Users"], _ = user.All()
	data["Categories"], _ = category.All()

    //generate template file
    allFiles := getTemplateFiles(tplFiles...)

    tmpl, err := template.New("").
        Funcs(template.FuncMap{
            "RouteName2URL": route.Name2URL,
        }).ParseFiles(allFiles...)
    logger.LogError(err)

    err = tmpl.ExecuteTemplate(c.Writer, name, data)
    logger.LogError(err)
}

func getTemplateFiles(tplFiles ...string) []string {
   
    viewDir := "resources/views/"

    
    for i, f := range tplFiles {
        tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
    }


    layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
    logger.LogError(err)
    return append(layoutFiles, tplFiles...)
}