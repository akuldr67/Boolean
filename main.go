package main

import (
	"fmt"
	"os"

	"github.com/akuldr67/Boolean/control"

	"github.com/akuldr67/Boolean/config"
)

func main() {
	fmt.Println("Boolean app!")

	err := config.ConnectDb()

	// config.DB.AutoMigrate(&models.Boolean{})

	if err != nil {
		fmt.Println("Unable to connect to database")
		os.Exit(1)
	}

	r := control.SetupRoutes()
	r.Run()
}
