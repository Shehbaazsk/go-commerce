package services

import (
	"errors"

	"github.com/shehbaazsk/go-commerce/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.DB.Create(user).Error
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := s.DB.Preload("UserProfile").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.DB.Preload("UserProfile").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) UpdateUser(id uint, updated *models.User) error {
	var user models.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return err
	}
	return s.DB.Model(&user).Updates(updated).Error
}

func (s *UserService) DeleteUser(id uint) error {
	return s.DB.Delete(&models.User{}, id).Error
}
