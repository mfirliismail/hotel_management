package controllers

import (
    "github.com/gin-gonic/gin"
    "hotel-management/models"
    "hotel-management/services"
    "net/http"
    "strconv"
)

type UserController struct {
    Service services.UserService
}

func (c *UserController) CreateUser(ctx *gin.Context) {
    var user models.User
    if err := ctx.BindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.Service.CreateUser(&user); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"user_id": user.ID})
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
    id := ctx.Param("id")

    userID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var user models.User
    if err := c.Service.GetUserByID(uint(userID), &user); err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    ctx.JSON(http.StatusOK, user)
}

// GetAllUsers retrieves all users.
func (c *UserController) GetAllUsers(ctx *gin.Context) {
    var users []models.User
    if err := c.Service.GetAllUsers(&users); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "status": "success",
        "data":   users,
    })
}