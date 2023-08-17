package database

import (
	"fmt"

	"github.com/nabil-y/zerogaspi-back/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDatabase() {
	DB, err = gorm.Open(sqlite.Open("zerogaspi.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection opened to database")
	DB.AutoMigrate(&model.User{}, &model.Perishable{})
	fmt.Println("Database auto-migrated")
}
