package main

import (
	"net/http"

	"example.com/event-booking-api/db"
	"example.com/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()

	server.GET("/getallevents", getAllEvents)
	server.POST("/createevent", saveEvent)

	server.Run(":8080")
}

func saveEvent(context *gin.Context) {

	var e models.Event
	err := context.ShouldBindJSON(&e)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	// Dummy id and user id generated - db will be handling this
	e.Id = 1
	e.UserId = 1
	sError := e.Save()
	if sError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "unable to create event. Try again later!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": e})
}

func getAllEvents(context *gin.Context) {
	e, getErr := models.GetAllEvents()
	if getErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve events. Try again later!"})
		return
	}
	context.JSON(http.StatusOK, e)
}
