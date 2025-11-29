package services

import (
	"errors"
	"yeni-proje-backend/internal/database"
	"yeni-proje-backend/internal/models"
)

// CreateAtolye creates a new workshop in the database.
func CreateAtolye(atolye *models.Atolye) (*models.Atolye, error) {
	if err := database.DB.Create(atolye).Error; err != nil {
		return nil, err
	}
	return GetAtolyeByID(atolye.ID)
}

// GetAllAtolyeler retrieves all workshops from the database.
func GetAllAtolyeler() ([]models.Atolye, error) {
	var atolyeler []models.Atolye
	if err := database.DB.Preload("Teacher").Preload("Category").Find(&atolyeler).Error; err != nil {
		return nil, err
	}
	return atolyeler, nil
}

// GetAtolyeByID retrieves a single workshop by its ID.
func GetAtolyeByID(id uint) (*models.Atolye, error) {
	var atolye models.Atolye
	if err := database.DB.Preload("Teacher").Preload("Category").First(&atolye, id).Error; err != nil {
		return nil, err
	}
	return &atolye, nil
}

// ParticipateInAtolye records a teacher's participation in a workshop.
func ParticipateInAtolye(atolyeID uint, teacherID uint) (*models.AtolyeParticipant, error) {
	_, err := GetAtolyeByID(atolyeID)
	if err != nil {
		return nil, errors.New("atölye bulunamadı")
	}

	participation := &models.AtolyeParticipant{
		AtolyeID:  atolyeID,
		TeacherID: teacherID,
	}

	if err := database.DB.Create(participation).Error; err != nil {
		return nil, errors.New("bu atölyeye zaten katıldınız veya bir hata oluştu")
	}

	var details models.AtolyeParticipant
	database.DB.Preload("Teacher").Preload("Atolye").First(&details, participation.ID)
	return &details, nil
}
