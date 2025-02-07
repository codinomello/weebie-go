package main

import (
	"os"
	"project/controllers"
	"project/db"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.Init()
	projectCollection := &controllers.GameController{
		Collection: &db.MongoCollection{Collection: database.Collection("games")},
	}

	r := gin.Default()

	// Configurar templates
	r.LoadHTMLGlob("views/templates/*")

	// Rotas
	//r.GET("/submit", gameCollection.SubmitFormHandler)
	r.POST("/submit", projectCollection.SubmitGameHandler)

	// Servir arquivos estáticos
	r.Static("/static", "./views/static")

	r.Run(":" + os.Getenv("8080"))
}
