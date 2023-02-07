package store

import (
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"fmt"
)

type Handler struct {
	DB *gorm.DB
}

var h = Handler{}

func NewDb() *gorm.DB {
	dbUri := fmt.Sprintf(`server=%s;user id=%s;password=%s;port=%s;database=%s`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))

	db, err := gorm.Open(sqlserver.Open(dbUri), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		log.Fatal(err)
	}
	h.DB = db
	return db
}

func Database() *gorm.DB {
	return h.DB
}
