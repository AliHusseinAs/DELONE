package models

import "time"

// Project represents the project table.
type Project struct {
	ID                     uint       `gorm:"primaryKey;column:projectId" json:"projectId"`
	TeacherID              uint       `gorm:"column:teacherId" json:"teacherId"`
	CategoryID             uint       `gorm:"column:categoryId" json:"categoryId"`
	ProjectName            string     `gorm:"column:projeAdi" json:"projeAdi"`
	Description            string     `gorm:"column:aciklama" json:"aciklama"`
	Text                   string     `gorm:"column:text" json:"text"`
	Slogan                 string     `gorm:"column:slogan" json:"slogan"`
	SubjectTag             string     `gorm:"column:konuEtiketi" json:"konuEtiketi"`
	StartDate              *time.Time `gorm:"column:baslangicTarihi" json:"baslangicTarihi"`
	EndDate                *time.Time `gorm:"column:bitisTarihi" json:"bitisTarihi"`
	EducationType          string     `gorm:"column:egitimTuru" json:"egitimTuru"`
	ParticipantLevel       string     `gorm:"column:katilimciDuzeyi" json:"katilimciDuzeyi"`
	QuotaInfo              string     `gorm:"column:kontenjanBilgisi" json:"kontenjanBilgisi"`
	ParticipationCondition string     `gorm:"column:katilimKosulu" json:"katilimKosulu"`
	Fee                    string     `gorm:"column:egitimUcreti" json:"egitimUcreti"`
	ContactPermission      string     `gorm:"column:iletisimOnay" json:"iletisimOnay"`
	PhotoPermission        string     `gorm:"column:fotoOnay" json:"fotoOnay"`

	// Relationships
	Teacher  Teacher  `gorm:"foreignKey:TeacherID" json:"teacher"`
	Category Category `gorm:"foreignKey:CategoryID" json:"category"`
}

func (Project) TableName() string {
	return "project"
}
