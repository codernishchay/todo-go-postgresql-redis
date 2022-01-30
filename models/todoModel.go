package models

import (
	"encoding/json"

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

func (i Todo) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(i)
	return bytes, err
}
