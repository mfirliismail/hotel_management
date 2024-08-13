// services/room_service.go
package services

import (
    "hotel-management/models"
    "gorm.io/gorm"
)

type RoomService interface {
    CreateRoom(room *models.Room) error
    GetRoomByID(id uint) (models.Room, error)
    UpdateRoom(id uint, room *models.Room) error
    DeleteRoom(id uint) error
    FilterRooms(category string, minPrice, maxPrice int) ([]models.Room, error)
}

type RoomServiceImpl struct {
    DB *gorm.DB
}

func (s *RoomServiceImpl) CreateRoom(room *models.Room) error {
    return s.DB.Create(room).Error
}

func (s *RoomServiceImpl) GetRoomByID(id uint) (models.Room, error) {
    var room models.Room
    err := s.DB.Preload("Reviews").First(&room, id).Error
    return room, err
}

func (s *RoomServiceImpl) UpdateRoom(id uint, room *models.Room) error {
    // Use Updates to only update fields that are non-zero in the struct
    return s.DB.Model(&models.Room{}).Where("id = ?", id).Updates(room).Error
}

func (s *RoomServiceImpl) DeleteRoom(id uint) error {
    return s.DB.Delete(&models.Room{}, id).Error
}

func (s *RoomServiceImpl) FilterRooms(category string, minPrice, maxPrice int) ([]models.Room, error) {
    var rooms []models.Room
    query := s.DB.Model(&models.Room{})
    if category != "" {
        query = query.Where("category = ?", category)
    }
    if minPrice > 0 {
        query = query.Where("price >= ?", minPrice)
    }
    if maxPrice > 0 {
        query = query.Where("price <= ?", maxPrice)
    }
    err := query.Find(&rooms).Error
    return rooms, err
}


