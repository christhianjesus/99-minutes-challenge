package entities

import (
	"gorm.io/gorm"
)

type Coordinate struct {
	address string
	zipcode string
	ext_num string
	int_num string
}

type Order struct {
	gorm.Model
	Status   string
	Quantity int
	Weight   int
}
