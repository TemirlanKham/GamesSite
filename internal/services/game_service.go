package services

import (
	"GamesSite/internal/models"
)

type GameRepository interface {
	GetAll() ([]models.Game, error)
	GetByID(id int) (*models.Game, error)
	Create(game *models.Game) error
	Update(id int, game *models.GameEdit) error
	Delete(gameID int) error
}

type GameService struct {
	repo GameRepository
}

func NewGameService(gameRepo GameRepository) *GameService {
	return &GameService{repo: gameRepo}
}
func (g *GameService) GetAllGames() ([]models.Game, error) {
	return g.repo.GetAll()
}
func (g *GameService) GetGameByID(id int) (*models.Game, error) {
	return g.repo.GetByID(id)
}
func (g *GameService) Create(name, description string, rating float64) (*models.Game, error) {
	game := &models.Game{
		Name:        name,
		Description: description,
		Rating:      rating,
	}
	err := g.repo.Create(game)
	return game, err
}

func (g *GameService) Update(id int, gameEdit *models.GameEdit) (*models.Game, error) {
	err := g.repo.Update(id, gameEdit)
	if err != nil {
		return nil, err
	}
	return g.GetGameByID(id)
}
func (g *GameService) DeleteGame(gameID int) error {
	return g.repo.Delete(gameID)
}
