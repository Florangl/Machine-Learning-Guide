package main

import (
	"log"

	"meditrack/configs"
	"meditrack/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	log.Println("Server running on http://localhost:8090")
	r.Run(":8090")
}
