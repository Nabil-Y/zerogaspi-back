package userRepository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nabil-y/zerogaspi-back/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]*model.User, error)
	GetUser(userId uuid.UUID) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(userId uuid.UUID, userUpdateRequest *model.User) (*model.User, error)
	DeleteUser(userId uuid.UUID) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) GetUsers() ([]*model.User, error) {
	var users []*model.User
	ur.DB.Find(&users)

	if err := ur.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetUser(userId uuid.UUID) (*model.User, error) {
	var user *model.User
	if err := ur.DB.Find(&user, userId).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) CreateUser(newUser *model.User) (*model.User, error) {
	if err := ur.DB.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}

func (ur *userRepository) UpdateUser(userId uuid.UUID, updatedUser *model.User) (*model.User, error) {
	var user *model.User
	if err := ur.DB.Find(&user, userId).Error; err != nil {
		return nil, err
	}
	if user.DeletedAt != nil || user.ID == uuid.Nil {
		return nil, errors.New("No User found")
	}
	if err := ur.DB.Model(&user).Updates(&updatedUser).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) DeleteUser(userId uuid.UUID) error {
	var user *model.User
	if err := ur.DB.First(&user, userId).Error; err != nil {
		return err
	}
	if user.DeletedAt != nil || user.ID == uuid.Nil {
		return errors.New("No User found")
	}

	if err := ur.DB.Delete(&user, userId).Error; err != nil {
		return err
	}
	return nil
}
