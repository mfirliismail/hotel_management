package services

import (
    "hotel-management/models"
    "gorm.io/gorm"
)

type PaymentService interface {
    CreatePayment(payment *models.Payment) error
    GetPaymentByID(paymentID string, payment *models.Payment) error
    UpdatePayment(payment *models.Payment) error
}

type PaymentServiceImpl struct {
    DB *gorm.DB
}

func (s *PaymentServiceImpl) CreatePayment(payment *models.Payment) error {
    return s.DB.Create(payment).Error
}

func (s *PaymentServiceImpl) GetPaymentByID(paymentID string, payment *models.Payment) error {
    return s.DB.Where("payment_id = ?", paymentID).First(payment).Error
}

func (s *PaymentServiceImpl) UpdatePayment(payment *models.Payment) error {
    return s.DB.Save(payment).Error
}
