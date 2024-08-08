package services

import (
    "hotel-management/models"
    "gorm.io/gorm"
)

type BookingService interface {
    CreateBooking(booking *models.Booking) error
    GetBookingByID(id uint) (models.Booking, error)
    UpdateBooking(booking *models.Booking) error
    DeleteBooking(id uint) error
}

type BookingServiceImpl struct {
    DB *gorm.DB
}

func (s *BookingServiceImpl) CreateBooking(booking *models.Booking) error {
    return s.DB.Create(booking).Error
}

func (s *BookingServiceImpl) GetBookingByID(id uint) (models.Booking, error) {
    var booking models.Booking
    err := s.DB.First(&booking, id).Error
    return booking, err
}

func (s *BookingServiceImpl) UpdateBooking(booking *models.Booking) error {
    return s.DB.Save(booking).Error
}

func (s *BookingServiceImpl) DeleteBooking(id uint) error {
    return s.DB.Delete(&models.Booking{}, id).Error
}
