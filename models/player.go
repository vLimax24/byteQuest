package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name string
	Class string
	PIN int
	Level int
	Experience int
	Bytes int
}
