package models

import (
	"errors"
	"gobasictinyurl/src/hashutility"
)

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) HashPassword(password string) {
	hashedPassword := hashutility.HashStr(password)
	user.Password = hashedPassword
}
func (user *User) CheckPassword(providedPassword string) error {
	compareResult := hashutility.CompareHashAndPassword(providedPassword, user.Password)
	if !compareResult {
		return errors.New("given password does not match for the current user")
	} else {
		return nil
	}
}
