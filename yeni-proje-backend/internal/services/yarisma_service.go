package services

import (
	"errors"
	"yeni-proje-backend/internal/database"
	"yeni-proje-backend/internal/models"
)

// CreateYarisma creates a new competition in the database.
func CreateYarisma(yarisma *models.Yarisma) (*models.Yarisma, error) {
	if err := database.DB.Create(yarisma).Error; err != nil {
		return nil, err
	}
	return GetYarismaByID(yarisma.ID)
}

// GetAllYarismalar retrieves all competitions from the database.
func GetAllYarismalar() ([]models.Yarisma, error) {
	var yarismalar []models.Yarisma
	if err := database.DB.Preload("Teacher").Preload("Category").Find(&yarismalar).Error; err != nil {
		return nil, err
	}
	return yarismalar, nil
}

// GetYarismaByID retrieves a single competition by its ID.
func GetYarismaByID(id uint) (*models.Yarisma, error) {
	var yarisma models.Yarisma
	if err := database.DB.Preload("Teacher").Preload("Category").First(&yarisma, id).Error; err != nil {
		return nil, err
	}
	return &yarisma, nil
}

// ParticipateInYarisma records a teacher's participation in a competition.
func ParticipateInYarisma(yarismaID uint, teacherID uint) (*models.YarismaParticipant, error) {
	_, err := GetYarismaByID(yarismaID)
	if err != nil {
		return nil, errors.New("yarışma bulunamadı")
	}

	participation := &models.YarismaParticipant{
		YarismaID: yarismaID,
		TeacherID: teacherID,
	}

	if err := database.DB.Create(participation).Error; err != nil {
		return nil, errors.New("bu yarışmaya zaten katıldınız veya bir hata oluştu")
	}

	var details models.YarismaParticipant
	database.DB.Preload("Teacher").Preload("Yarisma").First(&details, participation.ID)
	return &details, nil
}
