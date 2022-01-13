package database

import (
	"log"
	"os"
	"time"

	"github.com/dzonib/daily_sales_tracker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// export value
var DB *gorm.DB

func Connect() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable color
		},
	)

	dsn := "host=localhost user=postgres password=pass123 dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("Could not connect to Data Base")
	}

	DB = db

	// migration
	dbErr := db.AutoMigrate(&models.User{}, &models.Role{}, models.Permission{})

	if dbErr != nil {
		println(dbErr)
	}
}
