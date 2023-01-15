package orm

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string
	Password string
	Fname    string
	Lname    string
}
