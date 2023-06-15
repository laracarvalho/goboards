package schema

import (
	"gorm.io/gorm"
)

type Listing struct {
	gorm.Model
	Role string
	Company string
	Location string
	Remote bool
	Link string
	Salary int64
	Description string
}