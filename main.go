package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"richard-velito.github.io/voucher-app/controllers"
	"richard-velito.github.io/voucher-app/models"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Routes
	r.GET("/voucher", controllers.FindVouchers)
	r.GET("/voucher/:id", controllers.FindVoucher)
	r.POST("/voucher", controllers.CreateVoucher)
	r.PATCH("/voucher/:id", controllers.UpdateVoucher)
	r.DELETE("/voucher/:id", controllers.DeleteVoucher)

	r.Run(":4000")
}
