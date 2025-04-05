main
import (
"GamesSite/internal/models"
"GamesSite/internal/routes"
"github.com/gin-gonic/gin"
"gorm.io/driver/postgres"
"gorm.io/gorm"
"log"
)

func main() {
db, err := gorm.Open(postgres.Open("postgres://postgres:monkey@localhost:5432/gogamedb?sslmode=disable"), &gorm.Config{})

тут у меня юзер postgres
пароль monkey
порт на котором запущен
и название базы данных: gogamedb.
if err != nil {
log.Fatal("Error connecting to the database:", err)
}

	err = db.AutoMigrate(&models.Game{})
	if err != nil {
		log.Fatal("Error on migrating to the DB", err)
	}

	r := gin.Default()
	routes.SetupGames(r, db)
	r.Run(":8080")
}
Работает на порте 8080

Добавляем базу данных в main.

В папке delivery у меня идет обработка запроса, которая в случае ошибки выдаст ошибку.


type GameEdit struct {
Name        string  `json:"name"`
Description string  `json:"description"`
Rating      float64 `json:"rating"`
}
type Game struct {
Id          int     `json:"id"`
Name        string  `json:"name"`
Description string  `json:"description"`
Rating      float64 `json:"rating"`
}

Это главные структуры. 

Дальше в папке repository мы взаимодействуем с нашим БД.

В папке services просто все мои функций. 

и в папке routes находятся мои Энд поинты 

func SetupGames(r *gin.Engine, db *gorm.DB) {
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

пример в Postman
Заходим в POST. Вбиваем http://localhost:8080/api/games 
Потом выбираем Body -> raw -> JSON 

cкелет будет выглядит примерно так
{
"name": "RDR",
"description": "RPG",
"rating": 8.8
}

