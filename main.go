package main

import (
	"net/http"
	"strconv"

	"example.com/event-booking-api/db"
	"example.com/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()

	server.GET("/getallevents", getAllEvents)
	server.POST("/createevent", saveEvent)
	server.GET("/getevent/:id", getEventById)

	server.Run(":8080")
}

func saveEvent(ctx *gin.Context) {

	var e models.Event
	err := ctx.ShouldBindJSON(&e)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	// Dummy id and user id generated - db will be handling this
	e.Id = 1
	e.UserId = 1
	sError := e.Save()
	if sError != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "unable to create event. Try again later!"})
		return
	}
	ctx.JSON(http.StatusCreated, e)
}

func getAllEvents(ctx *gin.Context) {
	e, getErr := models.GetAllEvents()
	if getErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve events. Try again later!"})
		return
	}
	ctx.JSON(http.StatusOK, e)
}

func getEventById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, convErr := strconv.ParseInt(idStr, 10, 64)
	if convErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	e, getErr := models.GetEventById(int64(idInt))
	if getErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	ctx.JSON(http.StatusOK, *e)
}
