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

func (s *HotelServiceImpl) CreateHotel(hotel *models.Hotel) error {
    return s.DB.Create(hotel).Error
}

func (s *HotelServiceImpl) GetAllHotel(hotel *[]models.Hotel) error {
    return s.DB.Preload("Rooms").Find(hotel).Error
}
