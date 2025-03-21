package models

import (
	"fmt"
	"go-chat/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Datebase doesn't exist, creating now...")

		dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort)
		dbTemp, err := gorm.Open(mysql.Open(dsnWithoutDB), &gorm.Config{})
		if err != nil {
			fmt.Println("Failed to connect to MySQL:", err)
			return
		}

		createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;", cfg.DBName)
		if err := dbTemp.Exec(createDBSQL).Error; err != nil {
			fmt.Println("Database create failed:", err)
			return
		}

		fmt.Println("Database create success！")

		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Database reconnection failed:", err)
			return
		}

	}
	fmt.Println("Database connection successful:", cfg.DBHost+":"+cfg.DBPort)

	Migrate()
}

func Migrate() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		panic("Migration failed:" + err.Error())
	}
}
