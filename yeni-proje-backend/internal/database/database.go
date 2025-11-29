package database

import (
	"fmt"
	"log"

	"yeni-proje-backend/internal/config" // Kendi modül yolun

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB veritabanı bağlantısını kurar ve global DB değişkenine atar
func ConnectDB() {
	config.LoadConfig() // Önce config'i yükle

	user := config.GetEnv("DB_USER", "root")
	password := config.GetEnv("DB_PASSWORD", "bartukurnaz")
	host := config.GetEnv("DB_HOST", "127.0.0.1")
	port := config.GetEnv("DB_PORT", "3306")
	dbname := config.GetEnv("DB_NAME", "piri")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı:", err)
	}

	log.Println("Veritabanı bağlantısı başarılı!")
}
