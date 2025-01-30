package routes

import (
	"net/http"
	"strconv"

	"example.com/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

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
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	ctx.JSON(http.StatusOK, *e)
}

func updateEvent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, convErr := strconv.ParseInt(idStr, 10, 64)
	if convErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	_, getErr := models.GetEventById(idInt)
	if getErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	var updatedEvent models.Event
	bErr := ctx.ShouldBindJSON(&updatedEvent)
	if bErr != nil {
		ctx.JSON(http.StatusBadRequest, bErr.Error())
		return
	}

	updatedEvent.Id = idInt

	err := updatedEvent.Update()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated succesfully"})
}

func deleteEvent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	event, err := models.GetEventById(idInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": event})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, *event)
}
