package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home godoc
// @Summary Home
// @Tags common
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
