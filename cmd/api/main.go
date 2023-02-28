package main

import (
	"log"

	"github.com/edwbaeza/coverage-api/apps/coverage/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if godotenv.Load() != nil {
		log.Println("Error loading .env file")
	}
	engine := gin.Default()
	server.RegisterRouter(engine)
	log.Println("Starting...")
	log.Fatal(engine.Run(":3000"))
}
