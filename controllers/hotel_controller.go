package controllers

import (
    "github.com/gin-gonic/gin"
    "hotel-management/models"
    "hotel-management/services"
    "net/http"
)

type HotelController struct {
    Service services.HotelService
}

func (c *HotelController) CreateHotel(ctx *gin.Context) {
    var hotel models.Hotel
    if err := ctx.BindJSON(&hotel); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.Service.CreateHotel(&hotel); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": hotel.ID,
	})
}

// GetAllHotel retrieves all hotels.
func (c *HotelController) GetAllHotel(ctx *gin.Context) {
    var hotels []models.Hotel
    if err := c.Service.GetAllHotel(&hotels); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "status": "success",
        "data":   hotels,
    })
}