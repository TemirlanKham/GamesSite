package main

import (
	"GamesSite/internal/db"
	"GamesSite/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()
	routes.SetupGames(r, db.DB)
	r.Run(":8080")
}
