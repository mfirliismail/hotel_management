package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/midtrans/midtrans-go"
    "github.com/midtrans/midtrans-go/snap"
    "hotel-management/models"
    "hotel-management/services"
    "hotel-management/utils"
    "net/http"
)

type PaymentController struct {
    Service services.PaymentService
}

func (c *PaymentController) CreatePayment(ctx *gin.Context) {
    var payment models.Payment
    if err := ctx.BindJSON(&payment); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.Service.CreatePayment(&payment); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Use Midtrans to handle payment
    snapReq := &snap.Request{
        TransactionDetails: midtrans.TransactionDetails{
            OrderID:  payment.PaymentID,
            GrossAmt: int64(payment.Amount),
        },
        // Other fields...
    }

    snapResponse, err := utils.SnapGateway.CreateTransaction(snapReq)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"redirect_url": snapResponse.RedirectURL})
}

func (c *PaymentController) HandleNotification(ctx *gin.Context) {
    var notification map[string]interface{}
    if err := ctx.BindJSON(&notification); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Handle Midtrans notification
    paymentID := notification["order_id"].(string)
    transactionStatus := notification["transaction_status"].(string)

    var payment models.Payment
    // Assuming you have a method to access DB in service
    if err := c.Service.GetPaymentByID(paymentID, &payment); err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
        return
    }
    payment.Status = transactionStatus
    if err := c.Service.UpdatePayment(&payment); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "received"})
}
