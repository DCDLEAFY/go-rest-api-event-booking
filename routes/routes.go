package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/getallevents", getAllEvents)
	server.POST("/createevent", saveEvent)
	server.GET("/getevent/:id", getEventById)
	server.PUT("/updateevent/:id", updateEvent)
	server.DELETE("/deleteevent/:id", deleteEvent)
}
