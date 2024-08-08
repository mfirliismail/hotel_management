package controllers

import (
    "github.com/gin-gonic/gin"
    "hotel-management/models"
    "hotel-management/services"
    "net/http"
    "strconv"
)

type InvoiceController struct {
    Service services.InvoiceService
}

func (c *InvoiceController) CreateInvoice(ctx *gin.Context) {
    var invoice models.Invoice
    if err := ctx.BindJSON(&invoice); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.Service.CreateInvoice(&invoice); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"invoice_id": invoice.ID})
}

func (c *InvoiceController) GetInvoiceByID(ctx *gin.Context) {
    id := ctx.Param("id")

    invoiceID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
        return
    }

    var invoice models.Invoice
    if err := c.Service.GetInvoiceByID(uint(invoiceID), &invoice); err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
        return
    }

    ctx.JSON(http.StatusOK, invoice)
}
