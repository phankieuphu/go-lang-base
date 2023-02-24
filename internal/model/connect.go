package model

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnetDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		// panic("failed to connect database")
		return
	}
	database.AutoMigrate(&Voucher{})
	fmt.Println("Connect database successfully")
	DB = database
}
