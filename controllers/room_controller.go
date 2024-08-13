package controllers

import (
    "github.com/gin-gonic/gin"
    "hotel-management/models"
    "hotel-management/services"
    "net/http"
    "strconv"
)

type RoomController struct {
    Service services.RoomService
}

func (c *RoomController) CreateRoom(ctx *gin.Context) {
    var room models.Room
    if err := ctx.BindJSON(&room); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.Service.CreateRoom(&room); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"room_id": room.ID})
}

func (c *RoomController) GetRoomByID(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    room, err := c.Service.GetRoomByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "status": "success",
        "data": room,
    })
}

func (c *RoomController) UpdateRoom(ctx *gin.Context) {
    var room models.Room
    
    // Ambil ID dari URL path
    id, _ := strconv.Atoi(ctx.Param("id"))
    // Temukan room berdasarkan ID yang diberikan
    _, err := c.Service.GetRoomByID(uint(id));
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
        return
    }

    // Bind JSON data ke room model, kecuali ID
    if err := ctx.BindJSON(&room); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Simpan perubahan pada room
    if err := c.Service.UpdateRoom(uint(id), &room); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Room successfully updated",
        "data":    room,
    })
}


func (c *RoomController) DeleteRoom(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    if err := c.Service.DeleteRoom(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Room deleted"})
}

func (c *RoomController) FilterRooms(ctx *gin.Context) {
    category := ctx.Query("category")
    minPrice, _ := strconv.Atoi(ctx.Query("min_price"))
    maxPrice, _ := strconv.Atoi(ctx.Query("max_price"))

    rooms, err := c.Service.FilterRooms(category, minPrice, maxPrice)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "status": "success",
        "data": rooms,
    })
}
