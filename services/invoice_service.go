package services

import (
    "hotel-management/models"
    "gorm.io/gorm"
)

type InvoiceService interface {
    CreateInvoice(invoice *models.Invoice) error
    GetInvoiceByID(id uint, invoice *models.Invoice) error
}

type InvoiceServiceImpl struct {
    DB *gorm.DB
}

// CreateInvoice creates a new invoice in the database.
func (s *InvoiceServiceImpl) CreateInvoice(invoice *models.Invoice) error {
    return s.DB.Create(invoice).Error
}

// GetInvoiceByID retrieves an invoice by ID and stores it in the provided invoice pointer.
func (s *InvoiceServiceImpl) GetInvoiceByID(id uint, invoice *models.Invoice) error {
    return s.DB.First(invoice, id).Error
}
