package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// "gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
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
	_, d := os.LookupEnv("DOCKER")
	godotenv.Load(".env")
	dbConfig := dbConfig{
		Host:     "localhost",
		Port:     3306,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
	}
	if d {
		dbConfig.Host = "host.docker.internal"
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

	DB, err = gorm.Open("mysql", dbURL(buildDBConfig()))
	// DB, err = gorm.Open(mysql.Open(dbURL(buildDBConfig())), &gorm.Config{})

	if err != nil {
		return err
	}
	log.Println("Connected to database...")
	return nil
}
