package database

import (
	"fmt"
	"go_fiber/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBConn *gorm.DB
)

func Connect() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

func Insert(user *models.User) {
	db := DBConn
	db.Create(&user)
}
