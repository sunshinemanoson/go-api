package main

import (
	AuthController "golang/jwt-api/Controller/auth"
	"golang/jwt-api/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// binding from json
type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fname    string `json:"fname" binding:"required"`
	Lname    string `json:"lname" binding:"required"`
	// Avatar   string `json:"avatar"`
}
type Users struct {
	gorm.Model
	Username string
	Password string
	Fname    string
	Lname    string
}

func main() {
	orm.InitDB()

	// Migrate the schema
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)
	r.Run("localhost:8787")
}
