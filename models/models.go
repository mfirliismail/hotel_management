// models/models.go
package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Email    string `gorm:"unique"`
    Bookings []Booking
    Invoices []Invoice
}

type Hotel struct {
    gorm.Model
    Name       string
    Address    string
    Rooms      []Room
}

type Room struct {
    gorm.Model
    HotelID     uint   `gorm:"index"`
    Category    string
    Price       int
    Rating      float64
    Reviews     []Review
}

type Review struct {
    gorm.Model
    RoomID      uint   `gorm:"index"`
    UserID      uint   `gorm:"index"`
    Rating      float64
    Comment     string
}

type Booking struct {
    gorm.Model
    BookingID   string `gorm:"unique"`
    RoomID      uint   `gorm:"index"`
    UserID      uint   `gorm:"index"`
    CheckInDate string
    CheckOutDate string
    TotalAmount int
    PaymentID   string `gorm:"index"`
    InvoiceID   string `gorm:"index"`
    Payment     Payment
    Invoice     Invoice
    User        User
    Room        Room
}

type Payment struct {
    gorm.Model
    PaymentID   string `gorm:"unique"`
    BookingID   uint   `gorm:"index"`
    Amount      int
    Status      string
}

type Invoice struct {
    gorm.Model
    InvoiceID   string `gorm:"unique"`
    UserID      uint   `gorm:"index"`
    Amount      int
    Status      string
    User        User
}
