package repository

import (
	"GamesSite/internal/models"
	"gorm.io/gorm"
)

type GameRepositoryImpl struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) *GameRepositoryImpl {
	return &GameRepositoryImpl{db: db}
}

func (g *GameRepositoryImpl) GetAll() ([]models.Game, error) {
	var games []models.Game
	err := g.db.Find(&games).Error
	return games, err
}

func (g *GameRepositoryImpl) GetByID(id int) (*models.Game, error) {
	var game models.Game
	err := g.db.First(&game, id).Error
	return &game, err
}

func (g *GameRepositoryImpl) Create(game *models.Game) error {
	return g.db.Create(game).Error
}

func (g *GameRepositoryImpl) Update(id int, game *models.GameEdit) error {
	return g.db.Model(&models.Game{}).Where("id = ?", id).Omit("id, CreatedAt").Updates(game).Error
}

func (g *GameRepositoryImpl) Delete(gameID int) error {
	return g.db.Delete(&models.Game{}, gameID).Error
}
