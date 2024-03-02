package users_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterUserPayload struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func RegisterUser(c *gin.Context) {
	var newUser RegisterUserPayload
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
