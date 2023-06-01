package main

import (
	"E-Commerce/logger"
	"E-Commerce/middlewares"
	"E-Commerce/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.LogError(err, logger.GetFileName())
		panic(err)
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middlewares.Authentication())
	routes.AddressRoutes(router)
	routes.CartRoutes(router)

	port := os.Getenv("SERVICE_PORT")

	log.Fatal(router.Run(":" + port))
}
