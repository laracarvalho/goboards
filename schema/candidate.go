package schema

import (
	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	Name string
	Role string
	Company string
	Location string
	RemoteOnly bool
	LinkedIn string
	SalaryMin int64
	SalaryMax int64
	Description string
}