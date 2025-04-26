package routes

import (
	"GamesSite/internal/auth"
	"GamesSite/internal/delivery"
	"GamesSite/internal/middleware"
	"GamesSite/internal/repository"
	"GamesSite/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupGames(r *gin.Engine, db *gorm.DB) {
	gameRepo := repository.NewGameRepository(db)
	gameService := services.NewGameService(gameRepo)
	gameHandler := delivery.NewGameHandler(gameService)
	games := r.Group("api/games")
	{
		games.GET("/", gameHandler.GetAllGames)
		games.GET("/:id", gameHandler.GetGame)

		games.POST("/", middleware.RoleRequired("admin"), gameHandler.CreateGame)
		games.PUT("/:id", middleware.RoleRequired("admin"), gameHandler.UpdateGame)
		games.DELETE("/:id", middleware.RoleRequired("admin"), gameHandler.DeleteGame)
	}

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
	}

	protected := r.Group("api")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/me", auth.Me)
	}
}
