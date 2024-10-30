package cmd

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogPost struct {
	ID    int    `json:"id"` // unique identifier
	Title string `json:"title"`
	Body  string `json:"body"`
}

var db *sql.DB

func CreatePostHandler(c *gin.Context) {
	var post BlogPost

	// Bind the JSON to the struct
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}

	// Insert the post into the database
	query := "INSERT INTO posts (title, body) VALUES ($1, $2) RETURNING id"
	err := db.QueryRow(query, post.Title, post.Body).Scan(&post.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}

	// Return the post as a response
	c.JSON(http.StatusOK, gin.H {
		"message": "Post created successfully",
		"post": post,
	})
}