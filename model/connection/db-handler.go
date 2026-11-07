package connection

import (
	models "github/giovanifz/todolist-go/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Task{})
}
