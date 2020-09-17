package config

import (
	"fmt"

	// "github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type dbConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func buildDBConfig() *dbConfig {
	dbConfig := dbConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "boolean",
		Password: "booleanPw",
		DBName:   "boolean",
	}
	return &dbConfig
}

func dbURL(dbConfig *dbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func ConnectDb() error {
	var err error
	// db, err = gorm.Open("mysql", dbURL(buildDBConfig()))
	DB, err = gorm.Open(mysql.Open(dbURL(buildDBConfig())), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Connected to database...")
	return nil
}
