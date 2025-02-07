package controllers

import (
	"context"
	"net/http"
	"project/db"
	"project/models"
	"time"

	"github.com/gin-gonic/gin"
)

type GameController struct {
	Collection *db.MongoCollection
}

func (gc *GameController) SubmitGameHandler(c *gin.Context) {
	game := models.Game{
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		Author:      c.PostForm("author"),
		Executable:  c.PostForm("executable"),
		CoverImage:  c.PostForm("cover_image"),
		CreatedAt:   time.Now(),
	}

	_, err := gc.Collection.InsertOne(context.Background(), game)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

// Outros handlers...
