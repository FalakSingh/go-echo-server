package model

import (
	"echo-server/internal/utils"
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"-" bson:"password"`
	IsEmailVerified  bool   `json:"isEmailVerified" bson:"isEmailVerified"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}

func (u *User) CheckPass(givenPassword string) bool {
	return utils.ComparePassword(u.Password, givenPassword)
}
