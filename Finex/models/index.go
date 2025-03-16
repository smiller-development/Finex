package models

import {
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
}

type Config struct {
	Host	 string
	Port	 string
	User    string
	Password string
	DBName   string
}

var DB *gorm.DB

func InitDB(cfg Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	var err error
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
