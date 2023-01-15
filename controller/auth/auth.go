package auth

import (
	"golang/jwt-api/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fname    string `json:"fname" binding:"required"`
	Lname    string `json:"lname" binding:"required"`
	// Avatar   string `json:"avatar"`
}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userExist orm.Users
	orm.DB.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(200, gin.H{"status": "err", "message": "User Exist"})
		return
	}

	users := orm.Users{Username: json.Username, Password: json.Password, Fname: json.Fname, Lname: json.Lname}
	orm.DB.Create(&users)
	if users.ID > 0 {
		c.JSON(200, gin.H{"status": "ok", "message": "Registered.", "UserID": users.ID})
	} else {
		c.JSON(200, gin.H{"status": "err", "message": "Register Failed.", "UserID": users.ID})
	}
	c.JSON(200, gin.H{
		"message": json,
	})
}

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userExist orm.Users
	orm.DB.Where("username = ?", json.Username).First(&userExist)
	orm.DB.Where("password = ?", json.Password).First(&userExist)
	if userExist.ID == 0 {
		c.JSON(200, gin.H{"status": "err", "message": "User Not Exist"})
		return
	}

	// err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(json.Password))
	// if userExist.Password == 0 {
	// 	c.JSON(200, gin.H{"status": "err", "message": "User Not Exist"})
	// 	return
	// }
	// err := bcrypt.CompareAndPassword([]byte(userExist.Password), []byte(json.Password))

}
