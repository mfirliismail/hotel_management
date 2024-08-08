package services

import (
    "hotel-management/models"
    "gorm.io/gorm"
)

type UserService interface {
    CreateUser(user *models.User) error
    GetUserByID(id uint, user *models.User) error
    GetAllUsers(users *[]models.User) error // Tambahkan method ini
}

type UserServiceImpl struct {
    DB *gorm.DB
}

// CreateUser creates a new user in the database.
func (s *UserServiceImpl) CreateUser(user *models.User) error {
    return s.DB.Create(user).Error
}

// GetUserByID retrieves a user by ID and stores it in the provided user pointer.
func (s *UserServiceImpl) GetUserByID(id uint, user *models.User) error {
    return s.DB.First(user, id).Error
}

// GetAllUsers retrieves all users and stores them in the provided slice of users.
func (s *UserServiceImpl) GetAllUsers(users *[]models.User) error {
    return s.DB.Find(users).Error
}
