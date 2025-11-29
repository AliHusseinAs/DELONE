package models

import "time"

// Yarisma represents the yarisma table.
type Yarisma struct {
	ID                     uint       `gorm:"primaryKey;column:yarismaId" json:"yarismaId"`
	TeacherID              uint       `gorm:"column:teacherId" json:"teacherId"`
	CategoryID             uint       `gorm:"column:categoryId" json:"categoryId"`
	CompetitionName        string     `gorm:"column:atolyeAdi" json:"atolyeAdi"` // Şemada 'atolyeAdi' olarak kalmış.
	Description            string     `gorm:"column:aciklama" json:"aciklama"`
	Text                   string     `gorm:"column:baslik" json:"baslik"`
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

	// Relationships
	Teacher  Teacher  `gorm:"foreignKey:TeacherID" json:"teacher"`
	Category Category `gorm:"foreignKey:CategoryID" json:"category"`
}

func (Yarisma) TableName() string {
	return "yarisma"
}
