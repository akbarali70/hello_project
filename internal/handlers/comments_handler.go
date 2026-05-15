package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func GetComments(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/comments")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch comments",
		})
		return
	}
	defer resp.Body.Close()

	var comments []Comment

	if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to decode json",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":    len(comments),
		"comments": comments,
	})
}
