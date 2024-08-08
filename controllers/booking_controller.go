package controllers

import (
    "github.com/gin-gonic/gin"
    "hotel-management/models"
    "hotel-management/services"
    "net/http"
    "strconv"
)

type BookingController struct {
    Service services.BookingService
}

func (c *BookingController) CreateBooking(ctx *gin.Context) {
    var booking models.Booking
    if err := ctx.BindJSON(&booking); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate and process booking
    if err := c.Service.CreateBooking(&booking); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Generate invoice and handle payment here
    // ...

    ctx.JSON(http.StatusOK, gin.H{"booking_id": booking.ID})
}

func (c *BookingController) GetBookingByID(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    booking, err := c.Service.GetBookingByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, booking)
}

func (c *BookingController) UpdateBooking(ctx *gin.Context) {
    var booking models.Booking
    if err := ctx.BindJSON(&booking); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.Service.UpdateBooking(&booking); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, booking)
}

func (c *BookingController) DeleteBooking(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    if err := c.Service.DeleteBooking(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Booking deleted"})
}
