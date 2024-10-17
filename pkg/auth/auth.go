package auth

import (
	"errors"

	"github.com/Cynthia/goblog/app/models/user"
	"github.com/Cynthia/goblog/pkg/session"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func _getUID(c *gin.Context) string {
    _uid := session.Get(c,"uid")
    uid, ok := _uid.(string)
    if ok && len(uid) > 0 {
        return uid
    }
    return ""
}


func User(c *gin.Context) user.User {
    uid := _getUID(c)
    if len(uid) > 0 {
        _user, err := user.Get(uid)
        if err == nil {
            return _user
        }
    }
    return user.User{}
}

func Attempt(c *gin.Context,email string, password string) error {
    
    _user, err := user.GetByEmail(email)

  
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("账号不存在或密码错误")
        } else {
            return errors.New("内部错误")
        }
    }

    
    if !_user.ComparePassword(password) {
        return errors.New("账号不存在或密码错误")
    }

    
    session.Put(c,"uid", _user.GetStringID())

    return nil
}


func Login(c *gin.Context,_user user.User) {
    session.Put(c,"uid", _user.GetStringID())
}


func Logout(c *gin.Context) {
    session.Forget(c,"uid")
}


func Check(c *gin.Context) bool {
    return len(_getUID(c)) > 0
}