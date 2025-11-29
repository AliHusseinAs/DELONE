package models

import "time"

// YarismaParticipant represents a teacher's participation in a competition.
type YarismaParticipant struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	YarismaID uint      `gorm:"column:yarismaId;not null;uniqueIndex:idx_teacher_yarisma" json:"yarismaId"`
	TeacherID uint      `gorm:"column:teacherId;not null;uniqueIndex:idx_teacher_yarisma" json:"teacherId"`
	CreatedAt time.Time `json:"createdAt"`

	// Relationships for detailed responses
	Yarisma Yarisma `gorm:"foreignKey:YarismaID" json:"yarisma"`
	Teacher Teacher `gorm:"foreignKey:TeacherID" json:"teacher"`
}

func (YarismaParticipant) TableName() string {
	return "yarisma_participants"
}
