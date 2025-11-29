package models

import "time"

// ProjectParticipant represents a teacher's participation in a project.
type ProjectParticipant struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProjectID uint      `gorm:"column:projectId;not null;uniqueIndex:idx_teacher_project" json:"projectId"`
	TeacherID uint      `gorm:"column:teacherId;not null;uniqueIndex:idx_teacher_project" json:"teacherId"`
	CreatedAt time.Time `json:"createdAt"`

	// Relationships for detailed responses
	Project Project `gorm:"foreignKey:ProjectID" json:"project"`
	Teacher Teacher `gorm:"foreignKey:TeacherID" json:"teacher"`
}

func (ProjectParticipant) TableName() string {
	return "project_participants"
}
