package routes

import (
	"GamesSite/internal/delivery"
	"GamesSite/internal/repository"
	"GamesSite/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupGameRoutes(r *gin.Engine, db *gorm.DB) {
	gameRepo := repository.NewGameRepository(db)
	gameService := services.NewGameService(gameRepo)
	gameHandler := delivery.NewGameHandler(gameService)
	games := r.Group("api/games")
	{
		games.GET("/", gameHandler.GetAllGames)
		games.GET("/:id", gameHandler.GetGame)
		games.POST("/", gameHandler.CreateGame)
		games.PUT("/:id", gameHandler.UpdateGame)
		games.DELETE("/:id", gameHandler.DeleteGame)
	}
}
