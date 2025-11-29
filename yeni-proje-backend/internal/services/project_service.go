package services

import (
	"errors"
	"yeni-proje-backend/internal/database"
	"yeni-proje-backend/internal/models"
)

// CreateProject creates a new project in the database.
func CreateProject(project *models.Project) (*models.Project, error) {
	if err := database.DB.Create(project).Error; err != nil {
		return nil, err
	}
	// Return the project with preloaded data for a full response
	return GetProjectByID(project.ID)
}

// GetAllProjects retrieves all projects from the database.
func GetAllProjects() ([]models.Project, error) {
	var projects []models.Project
	// Preload related data for list view
	if err := database.DB.Preload("Teacher").Preload("Category").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

// GetProjectByID retrieves a single project by its ID.
func GetProjectByID(id uint) (*models.Project, error) {
	var project models.Project
	if err := database.DB.Preload("Teacher").Preload("Category").First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

// ParticipateInProject records a teacher's participation in a project.
func ParticipateInProject(projectID uint, teacherID uint) (*models.ProjectParticipant, error) {
	// Check if the project exists first
	_, err := GetProjectByID(projectID)
	if err != nil {
		return nil, errors.New("proje bulunamadı")
	}

	participation := &models.ProjectParticipant{
		ProjectID: projectID,
		TeacherID: teacherID,
	}

	// Create will automatically fail if the unique index constraint is violated
	if err := database.DB.Create(participation).Error; err != nil {
		// You might want to check for a specific duplicate key error here
		return nil, errors.New("bu projeye zaten katıldınız veya bir hata oluştu")
	}

	// Return the full participation details
	var details models.ProjectParticipant
	database.DB.Preload("Teacher").Preload("Project").First(&details, participation.ID)
	return &details, nil
}
