package routes

import (
	"cashapp/core"
	"cashapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(e *gin.Engine, s services.Services) {
	// CreateUser creates a new user account
	// @Summary Create a new user
	// @Description Create a new user account with a unique cash tag
	// @Tags users
	// @Accept json
	// @Produce json
	// @Param user body core.CreateUserRequest true "User details"
	// @Success 200 {object} core.Response "User created successfully"
	// @Failure 400 {object} map[string]string "Bad request"
	// @Failure 409 {object} map[string]string "Cash tag already taken"
	// @Failure 500 {object} map[string]string "Internal server error"
	// @Router /users [post]
	e.POST("/users", func(c *gin.Context) {

		var req core.CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		response := s.Users.CreateUser(req)

		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}

		c.JSON(response.Code, response.Meta)
	})

}
