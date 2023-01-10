package db

import (
	"fmt"

	"github.com/notblessy/go-listing/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB :nodoc:
func InitDB() *gorm.DB {
	err := config.LoadENV()
	if err != nil {
		logrus.Fatal(err)
	}

	dsn := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s`, config.DBUser(), config.DBPassword(), config.DBHost(), config.DBPort(), config.DBName())
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal(fmt.Sprintf("failed to connect: %s", err))
	}

	return db
}

// CloseDB :nodoc:
func CloseDB(db *gorm.DB) {
	postgres, err := db.DB()
	if err != nil {
		logrus.Fatal(fmt.Sprintf("failed to disconnect: %s", err))
	}

	err = postgres.Close()
	if err != nil {
		logrus.Fatal(fmt.Sprintf("failed to close: %s", err))
	}
}
