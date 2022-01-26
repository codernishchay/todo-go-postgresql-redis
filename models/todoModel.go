package models

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Todo struct {
	gorm.Model
	Todo        string
	Description string
	Done        bool
	Date        string
	Priority    string
}
