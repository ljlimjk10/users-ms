package users_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestController(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}
