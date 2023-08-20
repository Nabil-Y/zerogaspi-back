package userRepository

import (
	"github.com/nabil-y/zerogaspi-back/model"
)

type UserRepository interface {
	GetUsers() ([]*model.User, error)
	GetUser(userId string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(userId string, userUpdateRequest *model.User) (*model.User, error)
	DeleteUser(userId string) error
}
