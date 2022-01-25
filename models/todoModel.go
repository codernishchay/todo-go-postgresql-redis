package models

import (
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Todo struct {
	gorm.Model
	TodoID      uint
	What        string
	Description string
	Done        bool
	End         time.Time
	Start       time.Time
	Priority    int
	// Labels      []string
}
