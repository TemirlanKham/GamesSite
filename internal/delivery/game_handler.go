package delivery

import (
	"GamesSite/internal/models"
	"GamesSite/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewGameHandler(service *services.GameService) *GameHandler {
	return &GameHandler{service: service}
}

type GameHandler struct {
	service *services.GameService
}

func (h *GameHandler) GetAllGames(c *gin.Context) {
	games, _ := h.service.GetAllGames()
	c.JSON(http.StatusOK, games)
}

func (h *GameHandler) GetGame(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	game, err := h.service.GetGameByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, game)
}

func (h *GameHandler) CreateGame(c *gin.Context) {
	var gameCreate models.GameEdit
	if err := c.ShouldBindJSON(&gameCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newGame, err := h.service.Create(gameCreate.Name, gameCreate.Description, gameCreate.Rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newGame)
}
func (h *GameHandler) UpdateGame(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var gameEdit models.GameEdit
	if err := c.ShouldBindJSON(&gameEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedGame, err := h.service.Update(id, &gameEdit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedGame)
}

func (h *GameHandler) DeleteGame(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteGame(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Game Not Found"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Game Deleted"})
}
