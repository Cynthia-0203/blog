package user

import (
	"github.com/Cynthia/goblog/app/models"
	"github.com/Cynthia/goblog/pkg/password"
	"github.com/Cynthia/goblog/pkg/route"
)


type User struct {
    models.BaseModel

    Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
    Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
    Password string `gorm:"type:varchar(255)" valid:"password"`
    PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}


func (user *User) ComparePassword(_password string) bool {
    return password.CheckHash(_password, user.Password)
}


func (user User) Link() string {
    return route.Name2URL("users.show", "id", user.GetStringID())
}