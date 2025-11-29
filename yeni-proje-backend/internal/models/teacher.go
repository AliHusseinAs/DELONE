package models

import "time"

// Teacher represents the teacher table from the database schema.
type Teacher struct {
	ID        uint       `gorm:"primaryKey;column:teacherid" json:"teacherId"`
	BirthDate *time.Time `gorm:"column:birth_date" json:"birthDate,omitempty"`
	Email     string     `gorm:"column:email" json:"email"`
	GoogleID  string     `gorm:"column:google_id" json:"googleId,omitempty"`
	Name      string     `gorm:"column:name" json:"name"`
	Password  string     `gorm:"column:password" json:"-"` // Never expose password in JSON
	Surname   string     `gorm:"column:surname" json:"surname"`
}

// TableName specifies the table name for GORM.
func (Teacher) TableName() string {
	return "teacher"
}
