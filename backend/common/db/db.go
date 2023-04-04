package db

import (
	"fmt"
	"log"
	"os"
	"time"
	"zibon/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBAction int64

const (
	Create DBAction = iota
	Update
)

type QueryOptions struct {
	Count     bool
	Serialize bool
}

func Connect() *gorm.DB {
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")

	url := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, name)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// defer sqlDB.Close()

	return db
}

func AutoMigrate(DB *gorm.DB) error {
	models := []interface{}{
		models.UserModel{},
	}

	var err error
	for i := 0; i < len(models); i++ {
		mod := models[i]
		err = DB.AutoMigrate(&mod)
	}

	return err
}
