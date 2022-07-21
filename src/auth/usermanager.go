package auth

import (
	"errors"
	"gobasictinyurl/src/models"
)

type UserManager struct {
	users []*models.User
}

var instantiated *UserManager = nil

func SetupUserManager() {
	if instantiated == nil {
		instantiated = &UserManager{users: []*models.User{}}
	}
}

// func (store *UserManager) NewUserStore() *UserManager {
// 	return &UserManager{users: []*models.User{}}
// }

// Not implemented yet
func GetUserFromStorage(username string) (*models.User, error) {
	return nil, errors.New("not implemented yet")
}

func AddUserToStore(user *models.User) error {
	for _, v := range instantiated.users {
		if v.Username == user.Username {
			return errors.New("user already exists")
		}
	}

	instantiated.users = append(instantiated.users, user)

	return nil
}
