package models

import (
	"errors"
	"gobasictinyurl/src/helpers"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name       string `json:"name"`
	Username   string `json:"username" gorm:"unique"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password"`
	UrlEntries []UrlEntry
}

func (user *User) HashPassword(password string) {
	hashedPassword := helpers.HashStr(password)
	user.Password = hashedPassword
}
func (user *User) CheckPassword(providedPassword string) error {
	compareResult := helpers.CompareHashAndPassword(providedPassword, user.Password)
	if !compareResult {
		return errors.New("given password does not match for the current user")
	} else {
		return nil
	}
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewString()
	return
}
