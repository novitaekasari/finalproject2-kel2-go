package orderController

import (
	"finalprojec2-kel-go/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var orders []models.Order

	models.DB.Preload("Items").Find(&orders)
	c.JSON(http.StatusOK, gin.H{"orders" : orders})
}

func Show(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := models.DB.Preload("Items").First(&order, id).Error; err != nil {
		switch err {
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
				return
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
		}
	}

	c.JSON(http.StatusOK, gin.H{"order" : order})
}

func Create(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{"message": errorMessages})
		return
	}

	models.DB.Create(&order)
	c.JSON(http.StatusOK, gin.H{"order" : order})
}

func Update(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	models.DB.First(&order, id)

	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.Where("order_id = ?", id).Delete(models.Item{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus item"})
		return
	}

	if models.DB.Save(&order).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Data berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := models.DB.Where("order_id = ?", id).Delete(models.Item{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus item"})
		return
	}

	if models.DB.Delete(&order, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Data berhasil dihapus"})
}
