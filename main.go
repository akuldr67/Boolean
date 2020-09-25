package main

import (
	"log"
	"os"

	"github.com/akuldr67/Boolean/control"
	"github.com/akuldr67/Boolean/models"

	"github.com/akuldr67/Boolean/config"
)

func main() {
	// fmt.Println("Boolean app!")
	log.Println("Boolean App!")

	err := config.ConnectDb()

	config.DB.AutoMigrate(&models.Boolean{})

	if err != nil {
		log.Println("Unable to connect to database")
		os.Exit(1)
	}

	r := control.SetupRoutes()
	r.Run()
}
