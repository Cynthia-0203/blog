package controllers

import (
	"net/http"

	"github.com/Cynthia/goblog/app/models/user"
	"github.com/Cynthia/goblog/app/requests"
	"github.com/Cynthia/goblog/pkg/auth"
	"github.com/Cynthia/goblog/pkg/flash"
	"github.com/Cynthia/goblog/pkg/route"
	"github.com/Cynthia/goblog/pkg/view"
	"github.com/gin-gonic/gin"
)


type AuthController struct {
}

// type userForm struct {
//     Name            string `valid:"name"`
//     Email           string `valid:"email"`
//     Password        string `valid:"password"`
//     PasswordConfirm string `valid:"password_confirm"`
// }



func (*AuthController) Register(c *gin.Context) {
    view.RenderSimple(c, view.D{}, "auth.register")
}


func (*AuthController) DoRegister(c *gin.Context) {

    _user := user.User{
        Name:            c.PostForm("name"),
        Email:           c.PostForm("email"),
        Password:        c.PostForm("password"),
        PasswordConfirm: c.PostForm("password_confirm"),
    }


    errs := requests.ValidateRegistrationForm(_user)

    if len(errs) > 0 {
        view.RenderSimple(c, view.D{
            "Errors": errs,
            "User":   _user,
        }, "auth.register")
    } else {
        _user.Create()

        if _user.ID > 0 {
			flash.Success(c,"registered successfully...")
			auth.Login(c,_user)
            index:=route.Name2URL("articles.show", "id", _user.GetStringID())
            c.Redirect(http.StatusFound, index)
        } else {
            c.Writer.WriteHeader(http.StatusInternalServerError)
            c.String(http.StatusInternalServerError, "failed to register...")  
        }
    }
}

func (*AuthController) Login(c *gin.Context) {
    view.RenderSimple(c, view.D{}, "auth.login")
}


func (*AuthController) DoLogin(c *gin.Context) {
      
	  email := c.PostForm("email")
	  password := c.PostForm("password")
  
	  if err := auth.Attempt(c,email, password); err == nil {
		  flash.Success(c,"welcome back")
          c.Redirect(http.StatusFound, "/")
	  } else {
		  view.RenderSimple(c, view.D{
			  "Error":    err.Error(),
			  "Email":    email,
			  "Password": password,
		  }, "auth.login")
	  }
}

func (*AuthController) Logout(c *gin.Context) {
    auth.Logout(c)
	flash.Success(c,"logged out")
    c.Redirect(http.StatusFound, "/")
}