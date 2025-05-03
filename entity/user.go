package entity

import (
	"ticket-system/dto"
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"type:enum('admin','user');default:'user'" json:"role"`

	Tickets []Ticket `gorm:"foreignKey:UserID" json:"tickets,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) ToResponse() dto.AuthResponse {
	return dto.AuthResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}
