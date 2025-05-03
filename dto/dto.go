package dto

import "time"

// Request
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=admin user"`
}

// Login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Response
type AuthResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `json:"token,omitempty"`
}

// Ticket
type TicketRequest struct {
	EventID  uint `json:"event_id" binding:"required"`
	Quantity int  `json:"quantity" binding:"required,gt=0"`
}
