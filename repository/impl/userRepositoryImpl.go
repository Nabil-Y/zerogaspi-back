package userRepositoryImpl

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nabil-y/zerogaspi-back/model"
	"github.com/nabil-y/zerogaspi-back/repository"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository.UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}

func (ur *userRepositoryImpl) GetUsers() ([]*model.User, error) {
	var users []*model.User
	ur.DB.Find(&users)

	if err := ur.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepositoryImpl) GetUser(userId string) (*model.User, error) {
	if _, err := uuid.Parse(userId); err != nil {
		return nil, err
	}

	var user *model.User

	if err := ur.DB.First(&user, "id = ?", userId).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepositoryImpl) CreateUser(newUser *model.User) (*model.User, error) {
	if err := ur.DB.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}

func (ur *userRepositoryImpl) UpdateUser(userId string, updatedUser *model.User) (*model.User, error) {
	if _, err := uuid.Parse(userId); err != nil {
		return nil, err
	}

	var user *model.User

	if err := ur.DB.First(&user, "id = ?", userId).Error; err != nil {
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

func (ur *userRepositoryImpl) DeleteUser(userId string) error {
	if _, err := uuid.Parse(userId); err != nil {
		return err
	}

	var user *model.User
	if err := ur.DB.First(&user, "id = ?", userId).Error; err != nil {
		return err
	}
	if user.DeletedAt != nil || user.ID == uuid.Nil {
		return errors.New("No User found")
	}

	if err := ur.DB.Delete(&user, "id = ?", userId).Error; err != nil {
		return err
	}
	return nil
}
