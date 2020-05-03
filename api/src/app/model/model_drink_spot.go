package model

import (
	"github.com/jinzhu/gorm"
)

type DrinkSpot struct {
	gorm.Model
	Name        string
	Description string
	Latitude    float64
	Longitude   float64
	Quality     string
	UpVotes     int32
	DownVotes   int32
}
