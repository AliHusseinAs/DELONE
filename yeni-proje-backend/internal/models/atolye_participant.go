package models

import "time"

// AtolyeParticipant represents a teacher's participation in a workshop.
type AtolyeParticipant struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AtolyeID  uint      `gorm:"column:atolyeId;not null;uniqueIndex:idx_teacher_atolye" json:"atolyeId"`
	TeacherID uint      `gorm:"column:teacherId;not null;uniqueIndex:idx_teacher_atolye" json:"teacherId"`
	CreatedAt time.Time `json:"createdAt"`

	// Relationships for detailed responses
	Atolye  Atolye  `gorm:"foreignKey:AtolyeID" json:"atolye"`
	Teacher Teacher `gorm:"foreignKey:TeacherID" json:"teacher"`
}

func (AtolyeParticipant) TableName() string {
	return "atolye_participants"
}
