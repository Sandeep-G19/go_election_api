package main

import (
	"log"

	"go-rest-api/config"
	"go-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	routes.InitRoutes(r, db)

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
