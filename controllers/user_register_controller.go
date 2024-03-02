package users_controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/ljlimjk10/users-ms/auth"
	user_model "github.com/ljlimjk10/users-ms/models"
)

type RegisterUserPayload struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func RegisterUser(c *gin.Context, db *pg.DB) {
	var newUserPayload RegisterUserPayload
	err := c.ShouldBindJSON(&newUserPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPwd, err := auth.HashPassword(newUserPayload.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error. Please contact system admin."})
	}

	currentTime := time.Now()

	newUser := &user_model.User{
		Username:   newUserPayload.Username,
		Email:      newUserPayload.Email,
		HashedPass: hashedPwd,
		FirstName:  newUserPayload.FirstName,
		LastName:   newUserPayload.LastName,
		CreatedAt:  currentTime,
		ModifiedAt: currentTime,
	}

	if err := user_model.CreateUser(db, newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
