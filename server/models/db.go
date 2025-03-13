package models

import (
	"fmt"
	"go-chat/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Println("Database connection successful: ", cfg.DBHost+":"+cfg.DBPort)

	Migrate()
}

func Migrate() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		panic("Migration failed: " + err.Error())
	}
}
