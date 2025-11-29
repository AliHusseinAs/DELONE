package services

import (
	"yeni-proje-backend/internal/database"
	"yeni-proje-backend/internal/models"
)

// GetAllCategories veritabanındaki tüm etkinlik kategorilerini getirir
func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := database.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
