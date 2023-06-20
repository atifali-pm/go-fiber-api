package database

import (
	"log"
	"os"

	"githbub.com/atifali-pm/models"

	"github.com/jinzhu/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	DB *gorm.DB
}

var Database DbInstance

func ConnectDb() {

	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{
		Db: db,
	}

}
