package routes

import (
	"cashapp/core"
	"cashapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterPaymentRoutes registers payment-related routes
func RegisterPaymentRoutes(e *gin.Engine, s services.Services) {
	// SendMoney creates a new payment transaction
	// @Summary Create a payment transaction
	// @Description Send money from one user to another
	// @Tags payments
	// @Accept json
	// @Produce json
	// @Param payment body core.CreatePaymentRequest true "Payment details"
	// @Success 200 {object} core.Response "Payment processed successfully"
	// @Failure 400 {object} map[string]string "Bad request"
	// @Failure 500 {object} map[string]string "Internal server error"
	// @Router /payments [post]
	e.POST("/payments", func(c *gin.Context) {
		var req core.CreatePaymentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		response := s.Payments.SendMoney(req)
		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}

		c.JSON(response.Code, response.Meta)
	})
}
