package main

import (
	"fmt"
	"log"
	"sayeed1999/social-connect-golang-api/api/routes"
	"sayeed1999/social-connect-golang-api/config"
	"sayeed1999/social-connect-golang-api/infrastructure/cache"
	"sayeed1999/social-connect-golang-api/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()

	// Connect to database instance
	database := database.NewDatabase()
	if err := database.Connect(); err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	dbInstance := database.Instance()

	// Connect to cache (redis) instance
	cacheService := cache.NewCacheService()
	cacheService.Init()
	// defer cacheInstance.Close()

	// Initialize Gin engine
	app := gin.Default()

	// Initialize routes
	routes.InitRoutes(app, dbInstance, cacheService)

	addr := fmt.Sprintf("%v:%v", cfg.ListenIP, cfg.ListenPORT)
	log.Printf("%v api will listen on %v", cfg.API.NAME, addr)

	err := app.Run(addr)
	log.Fatal(err)
}
