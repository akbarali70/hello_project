package handlers

import (
	"context"
	"net/http"

	"hello_project/internal/db"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SeedUsers(c *gin.Context) {
	ctx := context.Background()

	rows := make([][]any, 0, 1_000_000)

	for i := 0; i < 1_000_000; i++ {
		rows = append(rows, []any{
			gofakeit.Name(),
			gofakeit.Email(),
			gofakeit.Number(18, 80),
		})
	}

	_, err := db.Pool.CopyFrom(
		ctx,
		pgx.Identifier{"users"},
		[]string{"name", "email", "age"},
		pgx.CopyFromRows(rows),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "1 million users inserted",
	})
}
