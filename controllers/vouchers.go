package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"richard-velito.github.io/voucher-app/models"
)

type CreateVoucherInput struct {
	Label       string `json:"label" binding:"required"`
	Description string `json:"description" binding:"required"`
	Important   bool   `json:"important"`
}

type UpdateVoucherInput struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	Important   bool   `json:"important"`
}

// GET /vouchers
// Find all vouchers
func FindVouchers(c *gin.Context) {
	var vouchers []models.Voucher
	models.DB.Find(&vouchers)

	c.JSON(http.StatusOK, gin.H{"vouchers": vouchers})
}

// GET /vouchers/:id
// Find a voucher
func FindVoucher(c *gin.Context) {
	// Get model if exist
	var voucher models.Voucher
	if err := models.DB.Where("id = ?", c.Param("id")).First(&voucher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"voucher": voucher})
}

// POST /vouchers
// Create new voucher
func CreateVoucher(c *gin.Context) {
	// Validate input
	var input CreateVoucherInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create voucher
	voucher := models.Voucher{
		Label: input.Label, Description: input.Description, Important: input.Important}
	models.DB.Create(&voucher)

	c.JSON(http.StatusOK, gin.H{"voucher": voucher})
}

// PATCH /vouchers/:id
// Update a voucher
func UpdateVoucher(c *gin.Context) {
	// Get model if exist
	var voucher models.Voucher
	if err := models.DB.Where("id = ?", c.Param("id")).First(&voucher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateVoucherInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&voucher).Updates(input)

	c.JSON(http.StatusOK, gin.H{"voucher": voucher})
}

// DELETE /vouchers/:id
// Delete a voucher
func DeleteVoucher(c *gin.Context) {
	// Get model if exist
	var voucher models.Voucher
	if err := models.DB.Where("id = ?", c.Param("id")).First(&voucher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&voucher)

	c.JSON(http.StatusOK, gin.H{"response": true})
}
