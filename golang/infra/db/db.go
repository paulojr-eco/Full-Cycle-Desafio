package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/paulojr-eco/full-cycle-desafio/domain/model"
	_ "gorm.io/driver/sqlite"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Fatalf("Error on loading .env file")
	}
}

func ConnectDB() (*gorm.DB, error) {
	var dsn string
	var db *gorm.DB
	var err error

	dsn = os.Getenv("dsn")

	db, err = gorm.Open(os.Getenv("dbType"), dsn)

	if err != nil {
		return nil, err
	}

	if os.Getenv("debug") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AutoMigrateDb") == "true" {
		db.AutoMigrate(&model.Product{})
	}

	return db, nil
}
