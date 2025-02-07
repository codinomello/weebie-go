package handlers

import (
	"context"
	"net/http"
	"project/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	// Initialize MongoDB client and collection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	collection = client.Database("projetodb").Collection("projetos")
}

func submitFormHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "submit.html", nil)
}

func submitGameHandler(c *gin.Context) {
	game := models.Game{
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		Author:      c.PostForm("author"),
		Executable:  c.PostForm("executable"),
		CoverImage:  c.PostForm("cover_image"),
		CreatedAt:   time.Now(),
	}

	_, err := collection.InsertOne(context.Background(), game)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func gameDetailsHandler(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var game models.Game
	err := collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&game)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.HTML(http.StatusOK, "game.html", gin.H{
		"game": game,
	})
}
