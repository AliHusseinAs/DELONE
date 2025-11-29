package handlers

import (
	"net/http"
	"yeni-proje-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// GetCategories veritabanındaki tüm etkinlik kategorilerini listeler
func GetCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kategoriler getirilemedi"})
		return
	}

	c.JSON(http.StatusOK, categories)
}
