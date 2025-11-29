package models

import "time"

// Atolye represents the atolye table.
type Atolye struct {
	ID                     uint       `gorm:"primaryKey;column:atolyeId"`
	TeacherID              uint       `gorm:"column:teacherId"`
	CategoryID             uint       `gorm:"column:categoryId"`
	WorkshopName           string     `gorm:"column:projeAdi"`
	Description            string     `gorm:"column:aciklama"`
	Text                   string     `gorm:"column:text"`
	Slogan                 string     `gorm:"column:slogan"`
	SubjectTag             string     `gorm:"column:konuEtiketi"` // Bu alan eksikti, ekledim
	StartDate              *time.Time `gorm:"column:baslangicTarihi"`
	EndDate                *time.Time `gorm:"column:bitisTarihi"`
	EducationType          string     `gorm:"column:egitimTuru"`
	ParticipantLevel       string     `gorm:"column:katilimciDuzeyi"`
	QuotaInfo              string     `gorm:"column:kontenjanBilgisi"`
	ParticipationCondition string     `gorm:"column:katilimKosulu"`
	Fee                    string     `gorm:"column:egitimUcreti"`
	ContactPermission      string     `gorm:"column:iletisimOnay"`
	PhotoPermission        string     `gorm:"column:fotoOnay"`

	// Relationships
	Teacher  Teacher  `gorm:"foreignKey:TeacherID"`
	Category Category `gorm:"foreignKey:CategoryID"`
}

func (Atolye) TableName() string {
	return "atolye"
}
