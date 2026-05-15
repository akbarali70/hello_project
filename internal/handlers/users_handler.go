package handlers

import (
	"context"
	"strconv"

	"hello_project/internal/db"
	"hello_project/internal/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	rows, err := db.Pool.Query(context.Background(), `
		SELECT id, name, email, age
		FROM users
		ORDER BY id DESC
	`)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User

	err := db.Pool.QueryRow(context.Background(), `
		SELECT id, name, email, age
		FROM users
		WHERE id = $1
	`, id).Scan(&user.ID, &user.Name, &user.Email, &user.Age)

	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := db.Pool.QueryRow(context.Background(), `
		INSERT INTO users (name, email, age)
		VALUES ($1, $2, $3)
		RETURNING id
	`, input.Name, input.Email, input.Age).Scan(&input.ID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, input)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Pool.Exec(context.Background(), `
		UPDATE users
		SET name = $1, email = $2, age = $3
		WHERE id = $4
	`, input.Name, input.Email, input.Age, id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	input.ID = id
	c.JSON(200, input)
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := db.Pool.Exec(context.Background(), `
		DELETE FROM users
		WHERE id = $1
	`, id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "user deleted",
	})
}
