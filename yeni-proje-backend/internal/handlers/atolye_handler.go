package handlers

import (
	"net/http"
	"strconv"
	"time"
	"yeni-proje-backend/internal/models"
	"yeni-proje-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// CreateAtolyeRequest defines the structure for incoming workshop creation JSON.
type CreateAtolyeRequest struct {
	CategoryID             uint       `json:"categoryId"`
	WorkshopName           string     `json:"projeAdi"`
	Description            string     `json:"aciklama"`
	Text                   string     `json:"text"`
	Slogan                 string     `json:"slogan"`
	SubjectTag             string     `json:"konuEtiketi"`
	StartDate              *time.Time `json:"baslangicTarihi"`
	EndDate                *time.Time `json:"bitisTarihi"`
	EducationType          string     `json:"egitimTuru"`
	ParticipantLevel       string     `json:"katilimciDuzeyi"`
	QuotaInfo              string     `json:"kontenjanBilgisi"`
	ParticipationCondition string     `json:"katilimKosulu"`
	Fee                    string     `json:"egitimUcreti"`
	ContactPermission      string     `json:"iletisimOnay"`
	PhotoPermission        string     `json:"fotoOnay"`
}

// CreateAtolye handles the API request to create a new workshop.
func CreateAtolye(c *gin.Context) {
	var req CreateAtolyeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veya eksik veri: " + err.Error()})
		return
	}

	teacherID := c.GetUint("teacherID")

	// Convert the request struct to the GORM model struct
	atolye := models.Atolye{
		TeacherID:              teacherID,
		CategoryID:             req.CategoryID,
		WorkshopName:           req.WorkshopName,
		Description:            req.Description,
		Text:                   req.Text,
		Slogan:                 req.Slogan,
		SubjectTag:             req.SubjectTag,
		StartDate:              req.StartDate,
		EndDate:                req.EndDate,
		EducationType:          req.EducationType,
		ParticipantLevel:       req.ParticipantLevel,
		QuotaInfo:              req.QuotaInfo,
		ParticipationCondition: req.ParticipationCondition,
		Fee:                    req.Fee,
		ContactPermission:      req.ContactPermission,
		PhotoPermission:        req.PhotoPermission,
	}

	createdAtolye, err := services.CreateAtolye(&atolye)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Atölye oluşturulamadı: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdAtolye)
}

// GetAtolyeler handles the API request to list all workshops.
func GetAtolyeler(c *gin.Context) {
	atolyeler, err := services.GetAllAtolyeler()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Atölyeler getirilemedi: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, atolyeler)
}

// GetAtolyeByID handles the API request to get a single workshop.
func GetAtolyeByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz atölye ID"})
		return
	}

	atolye, err := services.GetAtolyeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Atölye bulunamadı"})
		return
	}
	c.JSON(http.StatusOK, atolye)
}

// ParticipateInAtolye handles the API request to join a workshop.
func ParticipateInAtolye(c *gin.Context) {
	atolyeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz atölye ID"})
		return
	}

	teacherID := c.GetUint("teacherID")

	participation, err := services.ParticipateInAtolye(uint(atolyeID), teacherID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, participation)
}
