package handlers

import (
	"net/http"
	"strconv"
	"time"
	"yeni-proje-backend/internal/models"
	"yeni-proje-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// CreateProjectRequest defines the structure for incoming project creation JSON.
// JSON anahtarlarını veritabanı kolon isimleriyle aynı yapıyoruz.
type CreateProjectRequest struct {
	CategoryID             uint       `json:"categoryId"`
	ProjectName            string     `json:"projeAdi"`
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

// CreateProject handles the API request to create a new project.
func CreateProject(c *gin.Context) {
	// 1. Gelen JSON'u GORM modeline değil, request modeline bağla.
	var req CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veya eksik veri: " + err.Error()})
		return
	}

	// 2. Middleware'den teacherId'yi al.
	teacherID := c.GetUint("teacherID")

	// 3. Request verilerini GORM modeline dönüştür.
	project := models.Project{
		TeacherID:              teacherID,
		CategoryID:             req.CategoryID,
		ProjectName:            req.ProjectName,
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

	// 4. Servis katmanına GORM modelini gönder.
	createdProject, err := services.CreateProject(&project)
	if err != nil {
		// Servisten gelen spesifik hatayı döndür
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Proje oluşturulamadı: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProject)
}

// GetProjects handles the API request to list all projects.
func GetProjects(c *gin.Context) {
	projects, err := services.GetAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Projeler getirilemedi: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}

// GetProjectByID handles the API request to get a single project.
func GetProjectByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz proje ID"})
		return
	}

	project, err := services.GetProjectByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Proje bulunamadı"})
		return
	}
	c.JSON(http.StatusOK, project)
}

// ParticipateInProject handles the API request to join a project.
func ParticipateInProject(c *gin.Context) {
	projectID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz proje ID"})
		return
	}

	// The teacher who is participating is the one making the request
	teacherID := c.GetUint("teacherID")

	participation, err := services.ParticipateInProject(uint(projectID), teacherID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, participation)
}
