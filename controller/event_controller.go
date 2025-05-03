package controller

import (
	"net/http"
	"strconv"

	"ticket-system/entity"
	"ticket-system/service"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	service service.EventService
}

func NewEventController(service service.EventService) *EventController {
	return &EventController{service}
}

func (c *EventController) CreateEvent(ctx *gin.Context) {
	var event entity.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdEvent, err := c.service.CreateEvent(&event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdEvent)
}

func (c *EventController) GetAllEvents(ctx *gin.Context) {
	events, err := c.service.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func (c *EventController) GetEvent(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	event, err := c.service.GetEventByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (c *EventController) UpdateEvent(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var event entity.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = uint(id)
	updatedEvent, err := c.service.UpdateEvent(&event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedEvent)
}

func (c *EventController) DeleteEvent(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = c.service.DeleteEvent(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
