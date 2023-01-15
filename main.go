package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// binding from json
type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fname    string `json:"fname" binding:"required"`
	Lname    string `json:"lname" binding:"required"`
	// Avatar   string `json:"avatar"`
}

func main() {
	r := gin.Default()
	r.POST("/register", func(c *gin.Context) {
		var json Register
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": json,
		})
	})
	r.Run("localhost:8282")
}
