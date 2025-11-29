package handlers

import (
	"net/http"
	"strconv"
	"time"
	"yeni-proje-backend/internal/models"
	"yeni-proje-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// CreateYarismaRequest defines the structure for incoming competition creation JSON.
type CreateYarismaRequest struct {
	CategoryID             uint       `json:"categoryId"`
	CompetitionName        string     `json:"atolyeAdi"` // Şemaya göre bu kolon adı kullanılıyor
	Description            string     `json:"aciklama"`
	Text                   string     `json:"baslik"`
	Slogan                 string     `gorm:"column:slogan" json:"slogan"`
	SubjectTag             string     `gorm:"column:konuEtiketi" json:"konuEtiketi"`         // Düzeltildi
	StartDate              *time.Time `gorm:"column:baslangicTarihi" json:"baslangicTarihi"` // Pointer yapıldı
	EndDate                *time.Time `gorm:"column:bitisTarihi" json:"bitisTarihi"`         // Pointer yapıldı
	EducationType          string     `gorm:"column:egitimTuru" json:"egitimTuru"`
	ParticipantLevel       string     `gorm:"column:katilimciDuzeyi" json:"katilimciDuzeyi"`
	QuotaInfo              string     `gorm:"column:kontenjanBilgisi" json:"kontenjanBilgisi"`
	ParticipationCondition string     `gorm:"column:katilimKosulu" json:"katilimKosulu"`
	Fee                    string     `gorm:"column:egitimUcreti" json:"egitimUcreti"`
	Amount                 string     `gorm:"column:tutar" json:"tutar"`
	ContactPermission      string     `gorm:"column:iletisimOnay" json:"iletisimOnay"`
	PhotoPermission        string     `gorm:"column:fotoOnay" json:"fotoOnay"`
}

// CreateYarisma handles the API request to create a new competition.
func CreateYarisma(c *gin.Context) {
	// 1. Gelen JSON'u Request modeline bağla
	var req CreateYarismaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veya eksik veri: " + err.Error()})
		return
	}

	// 2. Middleware'den teacherId'yi al
	teacherID := c.GetUint("teacherID")

	// 3. Request verilerini GORM modeline dönüştür
	yarisma := models.Yarisma{
		TeacherID:       teacherID,
		CategoryID:      req.CategoryID,
		CompetitionName: req.CompetitionName,
		Description:     req.Description,
		Text:            req.Text,
	}

	// 4. Servis katmanına GORM modelini gönder
	createdYarisma, err := services.CreateYarisma(&yarisma)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Yarışma oluşturulamadı: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdYarisma)
}

// GetYarismalar handles the API request to list all competitions.
func GetYarismalar(c *gin.Context) {
	yarismalar, err := services.GetAllYarismalar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Yarışmalar getirilemedi: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, yarismalar)
}

// GetYarismaByID handles the API request to get a single competition.
func GetYarismaByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz yarışma ID"})
		return
	}

	yarisma, err := services.GetYarismaByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Yarışma bulunamadı"})
		return
	}
	c.JSON(http.StatusOK, yarisma)
}

// ParticipateInYarisma handles the API request to join a competition.
func ParticipateInYarisma(c *gin.Context) {
	yarismaID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz yarışma ID"})
		return
	}

	teacherID := c.GetUint("teacherID")

	participation, err := services.ParticipateInYarisma(uint(yarismaID), teacherID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, participation)
}
