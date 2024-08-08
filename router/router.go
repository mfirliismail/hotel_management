package router

import (
    "github.com/gin-gonic/gin"
    "hotel-management/controllers"
    "hotel-management/services"
    "hotel-management/utils"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Initialize services
    db := utils.InitDB()
    utils.InitMidtrans()

    userService := &services.UserServiceImpl{DB: db}
    invoiceService := &services.InvoiceServiceImpl{DB: db}
    roomService := &services.RoomServiceImpl{DB: db}
    bookingService := &services.BookingServiceImpl{DB: db}
    paymentService := &services.PaymentServiceImpl{DB: db}

    // Initialize controllers
    userController := &controllers.UserController{Service: userService}
    invoiceController := &controllers.InvoiceController{Service: invoiceService}
    roomController := &controllers.RoomController{Service: roomService}
    bookingController := &controllers.BookingController{Service: bookingService}
    paymentController := &controllers.PaymentController{Service: paymentService}

    // Define routes
    router.POST("/user", userController.CreateUser)
    router.GET("/user/:id", userController.GetUserByID)
    router.GET("/user", userController.GetAllUsers)

    router.POST("/invoice", invoiceController.CreateInvoice)
    router.GET("/invoice/:id", invoiceController.GetInvoiceByID)

    router.POST("/room", roomController.CreateRoom)
    router.GET("/room/:id", roomController.GetRoomByID)
    router.PUT("/room/:id", roomController.UpdateRoom)
    router.DELETE("/room/:id", roomController.DeleteRoom)
    router.GET("/rooms", roomController.FilterRooms)

    router.POST("/booking", bookingController.CreateBooking)
    router.GET("/booking/:id", bookingController.GetBookingByID)
    router.PUT("/booking/:id", bookingController.UpdateBooking)
    router.DELETE("/booking/:id", bookingController.DeleteBooking)

    router.POST("/payment", paymentController.CreatePayment)
    router.POST("/notification", paymentController.HandleNotification)

    return router
}
