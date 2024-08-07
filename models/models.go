package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Email    string `gorm:"unique"`
    Coins    []Coins
    Orders   []Order
    Invoices []Invoice `gorm:"foreignKey:UserID"` // Tambahkan foreign key
}

type Coins struct {
    gorm.Model
    UserID  uint `gorm:"index"`
    Balance int
}

type Order struct {
    gorm.Model
    OrderID     string `gorm:"unique"`
    UserID      uint   `gorm:"index"`
    Status      string
    TotalAmount int
    Invoices    []Invoice `gorm:"foreignKey:OrderID"` // Tambahkan foreign key
}

type Invoice struct {
    gorm.Model
    InvoiceID  string `gorm:"unique"`
    OrderID    uint   `gorm:"index"`
    UserID     uint   `gorm:"index"` // Tambahkan field UserID
    Amount     int
    Status     string
}

type TopUpTransaction struct {
    gorm.Model
    TransactionID string `gorm:"unique"`
    UserID        uint   `gorm:"index"`
    Amount        int
    Status        string
}
