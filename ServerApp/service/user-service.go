package service

import (
	"ServerApp/entity"
	"ServerApp/initializers"
	"errors"

	"gorm.io/gorm"
)

type UserService interface {
	Save(user entity.User) error
	FindAll() []entity.User
	FindOne(id string) (entity.User, error)
	Update(id string, user entity.User) (entity.User, error)
	Delete(id string) error
}

type userService struct {
	users []entity.User
}

func NewUserService() UserService {
	return &userService{}
}

func (service *userService) Save(user entity.User) error {
	err := initializers.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *userService) FindOne(id string) (entity.User, error) {
	var user entity.User
	result := initializers.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, errors.New("User not found")
		}
		return user, result.Error
	}

	return user, nil
}

func (service *userService) FindAll() []entity.User {
	var users []entity.User
	initializers.DB.Find(&users)
	return users
}

func (service *userService) Update(id string, user entity.User) (entity.User, error) {
	var userToUpdate entity.User
	err := initializers.DB.First(&userToUpdate, id).Error
	if err != nil {
		return userToUpdate, err
	}

	err = initializers.DB.Model(&userToUpdate).Updates(user).Error
	if err != nil {
		return userToUpdate, err
	}

	return userToUpdate, nil
}

func (service *userService) Delete(id string) error {
	var user entity.User
	result := initializers.DB.Delete(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
