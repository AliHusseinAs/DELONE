package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"yeni-proje-backend/internal/config"
	"yeni-proje-backend/internal/database"
	"yeni-proje-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware JWT token'ı doğrular ve e-posta adresinden öğretmeni bulur.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header gerekli"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token formatı hatalı"})
			return
		}

		secretKey := config.GetEnv("JWT_SECRET_KEY", "4de5e3c7-00ef-4d34-a183-e423a31c3e17") // .env'den okuması için

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("beklenmeyen imzalama metodu: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Geçersiz veya süresi dolmuş token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token claims okunamadı"})
			return
		}

		// --- DEĞİŞİKLİK BURADA BAŞLIYOR ---

		// 1. Token'dan 'sub' alanındaki e-posta adresini oku
		teacherEmail, ok := claims["sub"].(string)
		if !ok || teacherEmail == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token içinde geçerli bir kullanıcı (sub) bulunamadı"})
			return
		}

		// 2. E-posta adresini kullanarak veritabanından öğretmeni bul
		var teacher models.Teacher
		if err := database.DB.Where("email = ?", teacherEmail).First(&teacher).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token'a ait öğretmen veritabanında bulunamadı"})
			return
		}

		// 3. Öğretmenin ID'sini ve e-postasını context'e ekle
		// Sonraki handler'lar bu bilgilere erişebilecek.
		c.Set("teacherID", teacher.ID)
		c.Set("teacherEmail", teacher.Email)

		c.Next()
	}
}
