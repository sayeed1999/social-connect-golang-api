package main

import (
	"fmt"
	"log"
	"sayeed1999/social-connect-golang-api/api/routes"
	"sayeed1999/social-connect-golang-api/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()

	// Connect to database instance
	// database.Connect()

	// Initialize Gin engine
	app := gin.Default()

	// Initialize routes
	routes.InitRoutes(app)

	addr := fmt.Sprintf("%v:%v", cfg.ListenIP, cfg.ListenPort)
	log.Printf("%v api will listen on %v", cfg.API.Name, addr)

	err := app.Run(addr)
	log.Fatal(err)
}
