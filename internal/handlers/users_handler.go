package handlers

import (
	"errors"

	"hello_project/internal/db"
	"hello_project/internal/dto"
	"hello_project/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUsers godoc
// @Summary Get users
// @Description Get latest 100 users
// @Tags users
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User

	if err := db.DB.Order("id desc").Limit(100).Find(&users).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

// GetUserByID godoc
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var input dto.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}

	if err := db.DB.Create(&user).Error; err != nil {

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(409, gin.H{
				"error": "email already exists",
			})
			return
		}

		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, user)
}

// UpdateUser godoc
// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User data"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "user not found",
		})
		return
	}

	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email

	if err := db.DB.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

// DeleteUser godoc
// @Summary Delete user
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := db.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "user deleted",
	})
}
