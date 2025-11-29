package main

import (
	"net/http"
	"yeni-proje-backend/internal/database"
	"yeni-proje-backend/internal/handlers"
	"yeni-proje-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	router := gin.Default()

	api := router.Group("/api")
	{
		// --- Project Routes ---
		projectRoutes := api.Group("/projects")
		{
			// Public routes (No JWT needed)
			projectRoutes.GET("", handlers.GetProjects)
			projectRoutes.GET("/:id", handlers.GetProjectByID)

			// Protected routes (JWT required)
			projectRoutes.POST("", middleware.AuthMiddleware(), handlers.CreateProject)
			projectRoutes.POST("/:id/participate", middleware.AuthMiddleware(), handlers.ParticipateInProject)
		}

		// --- Atolye (Workshop) Routes ---
		atolyeRoutes := api.Group("/workshops")
		{
			// Public routes
			atolyeRoutes.GET("", handlers.GetAtolyeler)
			atolyeRoutes.GET("/:id", handlers.GetAtolyeByID)

			// Protected routes
			atolyeRoutes.POST("", middleware.AuthMiddleware(), handlers.CreateAtolye)
			atolyeRoutes.POST("/:id/participate", middleware.AuthMiddleware(), handlers.ParticipateInAtolye)
		}

		// --- Yarisma (Competition) Routes ---
		yarismaRoutes := api.Group("/competitions") // URL'i İngilizce yapmak daha standart
		{
			// Public routes
			yarismaRoutes.GET("", handlers.GetYarismalar)
			yarismaRoutes.GET("/:id", handlers.GetYarismaByID)

			// Protected routes
			yarismaRoutes.POST("", middleware.AuthMiddleware(), handlers.CreateYarisma)
			yarismaRoutes.POST("/:id/participate", middleware.AuthMiddleware(), handlers.ParticipateInYarisma)
		}

		// Buraya daha sonra meta, profil gibi diğer rotalar eklenebilir.
		// Örneğin: api.GET("/categories", handlers.GetCategories)

		// Test için Ping Rotası
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
	}

	router.Run(":8080")
}
