package auth

import (
	"errors"
	"gobasictinyurl/src/models"
	"gobasictinyurl/src/persistence"
)

func GetUserFromStorageByEmail(email string) (*models.User, error) {
	var user models.User

	record := persistence.Instance.Where("email = ?", email).First(&user)

	if record.Error != nil {
		return nil, errors.New("user not found")
	} else {
		return &user, nil
	}
}

func GetUserFromStorageEagerById(userid string) (*models.User, error) {
	var user models.User

	record := persistence.Instance.Preload("UrlEntries").Where("id = ?", userid).First(&user)

	if record.Error != nil {
		return nil, errors.New("user not found")
	} else {
		return &user, nil
	}
}

func AddUserToStore(user *models.User) error {
	record := persistence.Instance.Create(&user)

	return record.Error
}
