package main

import (
	"gorm.io/gorm"
	"github.com/folklore/task-5-vix-btpns-ario/router"
	"github.com/folklore/task-5-vix-btpns-ario/database"
)

var db *gorm.DB = database.OpenConnection()

func main() {
	defer database.CloseConnection(db)

	router.AuthRouter()
	router.PhotoRouter()
	router.UserRouter()
}