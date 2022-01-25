package models

import (
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	UserID    uint
	Username  string
	FirstName string
	LastName  string
	JoinedOn  time.Time
	Password  string
	Todos     []Todo `gorm:"foreignKey:TodoID"`
	Email     string
}
