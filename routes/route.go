package route

import (
	"ticket-system/config"
	"ticket-system/controller"
	"ticket-system/middleware"
	"ticket-system/repository"
	"ticket-system/service"

	"github.com/gin-gonic/gin"
)

func ConnectRoutes(r *gin.Engine) {

	userRepo := repository.NewAuthRepository(config.DB)
	ticketRepo := repository.NewTicketRepository(config.DB)
	eventRepo := repository.NewEventRepository(config.DB)
	reportRepo := repository.NewReportRepository(config.DB)

	userService := service.NewAuthService(userRepo)
	ticketService := service.NewTicketService(ticketRepo, eventRepo)
	eventService := service.NewEventService(eventRepo)
	reportService := service.NewReportService(reportRepo)

	userController := controller.NewAuthController(userService)
	ticketController := controller.NewTicketController(ticketService)
	eventController := controller.NewEventController(eventService)
	reportController := controller.NewReportController(reportService)

	// authController *controller.UserController,
	// eventController *controller.EventController,
	// ticketController *controller.TicketController,
	// reportController *controller.ReportController,

	api := r.Group("/api")

	// Auth routes (public)
	api.POST("/register", userController.Register)
	api.POST("/login", userController.Login)

	// Event routes
	event := api.Group("/events")
	{
		// Public routes
		event.GET("", eventController.GetAllEvents)
		event.GET("/:id", eventController.GetEvent)

		// Protected admin routes
		event.Use(middleware.AuthMiddleware(), middleware.RoleAuthorization("admin"))
		{
			event.POST("", eventController.CreateEvent)
			event.PUT("/:id", eventController.UpdateEvent)
			event.DELETE("/:id", eventController.DeleteEvent)
		}
	}

	// Ticket routes
	ticket := api.Group("/tickets")
	{
		// Semua endpoint tiket sekarang membutuhkan autentikasi
		ticket.Use(middleware.AuthMiddleware())
		{
			// Endpoint pembelian tiket (untuk user yang login)
			ticket.POST("", ticketController.PurchaseTicket)

			// User melihat tiket sendiri
			ticket.GET("/my-tickets", ticketController.GetMyTickets)

			// User membatalkan tiket
			ticket.PATCH("/:id/cancel", ticketController.CancelTicket)

			// Admin-only routes
			adminTicket := ticket.Group("")
			adminTicket.Use(middleware.RoleAuthorization("admin"))
			{
				adminTicket.GET("/all", ticketController.GetAllTickets)                // Admin melihat semua tiket
				adminTicket.GET("/user/:user_id", ticketController.GetTicketsByUserID) // Admin melihat tiket by user_id
			}
		}
	}

	// Report routes (Admin only)
	report := api.Group("/reports")
	report.Use(middleware.AuthMiddleware(), middleware.RoleAuthorization("admin"))
	{
		report.GET("/summary", reportController.SummaryReport)
		report.GET("/event/:id", reportController.EventReport)
	}

	// Optional: Add health check endpoint
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
