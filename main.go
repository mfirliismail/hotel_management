package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
    "github.com/midtrans/midtrans-go"
    "github.com/midtrans/midtrans-go/snap"
    "log"
    "net/http"
    "os"
    "payment-go/models"
)

var (
    db          *gorm.DB
    snapGateway *snap.Client
)

func init() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}

func main() {
    // Setup database connection
    dsn := os.Getenv("DATABASE_URL")
    var err error
    db, err = gorm.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Migrate the schema
    db.AutoMigrate(&models.User{}, &models.Coins{}, &models.Order{}, &models.Invoice{}, &models.TopUpTransaction{})

// Get Midtrans keys from environment variables
serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
clientKey := os.Getenv("MIDTRANS_CLIENT_KEY")

if serverKey == "" || clientKey == "" {
    log.Fatal("MIDTRANS_SERVER_KEY or MIDTRANS_CLIENT_KEY not set in environment variables")
}

// Setup Midtrans Snap client
snapGateway = &snap.Client{}
snapGateway.New(serverKey, midtrans.Sandbox) 

    // Setup Gin router
    router := gin.Default()
    router.POST("/order", createOrder)
    router.POST("/invoice", createInvoice)
    router.POST("/payment", createPayment)
    router.POST("/notification", handleNotification)

    router.Run(":8080")
}

func createOrder(c *gin.Context) {
    var order models.Order
    if err := c.BindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save order to database
    db.Create(&order)

    c.JSON(http.StatusOK, gin.H{"order_id": order.OrderID})
}

func createInvoice(c *gin.Context) {
    var invoice models.Invoice
    if err := c.BindJSON(&invoice); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save invoice to database
    db.Create(&invoice)

    c.JSON(http.StatusOK, gin.H{"invoice_id": invoice.InvoiceID})
}

func createPayment(c *gin.Context) {
    var payment models.TopUpTransaction
    if err := c.BindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create Midtrans transaction
   // Initiate Customer address
	custAddress := &midtrans.CustomerAddress{
		FName:       "John",
		LName:       "Doe",
		Phone:       "081234567890",
		Address:     "Baker Street 97th",
		City:        "Jakarta",
		Postcode:    "16000",
		CountryCode: "IDN",
	}

	// Initiate Snap Request
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
            OrderID:  payment.TransactionID,
            GrossAmt: int64(payment.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    "John",
			LName:    "Doe",
			Email:    "john@doe.com",
			Phone:    "081234567890",
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "ITEM1",
				Price: int64(payment.Amount),
				Qty:   1,
				Name:  "Someitem",
			},
		},
	}

    snapResponse, err := snapGateway.CreateTransaction(snapReq)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Save payment to database
    db.Create(&payment)

    c.JSON(http.StatusOK, gin.H{"redirect_url": snapResponse.RedirectURL})
}


func handleNotification(c *gin.Context) {
    var notification map[string]interface{}
    if err := c.BindJSON(&notification); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Extract order ID and transaction status from the notification
    orderID := notification["order_id"].(string)
    transactionStatus := notification["transaction_status"].(string)

    // Handle Midtrans notification
    var payment models.TopUpTransaction
    db.Where("transaction_id = ?", orderID).First(&payment)
    payment.Status = transactionStatus
    db.Save(&payment)

    // Update invoice status
    var invoice models.Invoice
    db.Where("id = ?", payment.UserID).First(&invoice)
    invoice.Status = "paid"
    db.Save(&invoice)

    c.JSON(http.StatusOK, gin.H{"status": "received"})
}
