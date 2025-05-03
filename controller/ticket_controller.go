// controller/ticket_controller.go
package controller

import (
	"net/http"
	"strconv"

	"ticket-system/dto"
	"ticket-system/entity"
	"ticket-system/service"

	"github.com/gin-gonic/gin"
)

type TicketController interface {
	PurchaseTicket(ctx *gin.Context)
	GetMyTickets(ctx *gin.Context)
	GetAllTickets(ctx *gin.Context)
	GetTicketsByUserID(ctx *gin.Context)
	CancelTicket(ctx *gin.Context)
}

type ticketController struct {
	ticketService service.TicketService
}

func NewTicketController(ticketService service.TicketService) TicketController {
	return &ticketController{
		ticketService: ticketService,
	}
}

func (c *ticketController) PurchaseTicket(ctx *gin.Context) {
	// Dapatkan userID dari JWT
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var request dto.TicketRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pindahkan data dari DTO ke Entity
	ticket := entity.Ticket{
		UserID:   userID.(uint),
		EventID:  request.EventID,
		Quantity: request.Quantity,
		Status:   "tersedia", // atau status awal lainnya
	}

	result, err := c.ticketService.PurchaseTicket(&ticket)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
func (c *ticketController) GetMyTickets(ctx *gin.Context) {
	// Get user ID from context (set by AuthMiddleware)
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	tickets, err := c.ticketService.GetMyTickets(userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tickets)
}

func (c *ticketController) GetAllTickets(ctx *gin.Context) {
	// Check if user is admin
	userRole, exists := ctx.Get("user_role")
	if !exists || userRole.(string) != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	tickets, err := c.ticketService.GetAllTickets()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tickets)
}

func (c *ticketController) GetTicketsByUserID(ctx *gin.Context) {
	// Check if user is admin
	userRole, exists := ctx.Get("user_role")
	if !exists || userRole.(string) != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID, err := strconv.ParseUint(ctx.Param("user_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	tickets, err := c.ticketService.GetTicketsByUserID(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tickets)
}

func (c *ticketController) CancelTicket(ctx *gin.Context) {
	// Get user info from context (set by AuthMiddleware)
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	isAdmin := false
	if userRole, exists := ctx.Get("user_role"); exists && userRole.(string) == "admin" {
		isAdmin = true
	}

	ticketID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ticket ID"})
		return
	}

	err = c.ticketService.CancelTicket(uint(ticketID), userID.(uint), isAdmin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "ticket cancelled successfully"})
}
