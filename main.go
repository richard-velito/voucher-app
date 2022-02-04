package main

import (
	"github.com/gin-gonic/gin"
	"richard-velito.github.io/voucher-app/controllers"
	"richard-velito.github.io/voucher-app/models"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/voucher", controllers.FindVouchers)
	r.GET("/voucher/:id", controllers.FindVoucher)
	r.POST("/voucher", controllers.CreateVoucher)
	r.PATCH("/voucher/:id", controllers.UpdateVoucher)
	r.DELETE("/voucher/:id", controllers.DeleteVoucher)

	r.Run(":4000")
}
