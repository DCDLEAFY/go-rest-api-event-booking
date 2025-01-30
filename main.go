package main

import (
	"net/http"
	"strconv"

	"example.com/event-booking-api/db"
	"example.com/event-booking-api/models"
	"example.com/event-booking-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
