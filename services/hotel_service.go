package services

import (
    "hotel-management/models"
    "gorm.io/gorm"
)

type HotelService interface {
    CreateHotel(hotel *models.Hotel) error
	GetAllHotel(hotel *[]models.Hotel) error
}

type HotelServiceImpl struct {
    DB *gorm.DB
}

func (s *HotelServiceImpl) GetAllHotel(hotel *[]models.User) error {
    return s.DB.Find(hotel).Error
}
