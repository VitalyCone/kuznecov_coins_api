package dtos

import (
	"strings"

	"github.com/VitalyCone/account/internal/app/model"
)

type CreateUserDto struct {
	Username   string `json:"username" form:"username" validate:"required,alphanum,min=3,max=32"`
	Password   string `json:"password" form:"password" validate:"required,min=3,max=32"`
	FirstName  string `json:"first_name" form:"first_name" validate:"required,max=50"`
	SecondName string `json:"second_name" form:"second_name" validate:"required,max=50"`
}
func (c *CreateUserDto) ToModel(passHash string) model.User {
	return model.User{
		Username:     strings.ToLower(c.Username),
		PasswordHash: passHash,
		FirstName:    c.FirstName,
		SecondName:   c.SecondName,
	}
}

type ModifyUserDto struct {
	Username    string `json:"username" form:"username" validate:"required,alphanum,min=3,max=32"`
	OldPassword string `json:"old_password" form:"old_password" validate:"required,min=3,max=32"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required,min=3,max=32"`
	Avatar      string `json:"avatar" form:"avatar"`
	FirstName   string `json:"first_name" form:"first_name" validate:"required,max=50"`
	SecondName  string `json:"second_name" form:"second_name" validate:"required,max=50"`
}
func (u *ModifyUserDto) ToModel(passHash string) model.User {
	return model.User{}
}

type UserDto struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

