package main

import (
	"byteQuest/cmd"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Player struct {
	Name string
	Class string
	Level int
	Experience int
	Tools []string
	CurrentLevel int
}

func initDB() {
	var err error

	db, err := gorm.Open(sqlite.Open("players.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate()
}

func main() {
	initDB()

	cmd.Execute()
}