package entity

import "time"

type EventStatus string

const (
	StatusActive    EventStatus = "Aktif"
	StatusOngoing   EventStatus = "Berlangsung"
	StatusCompleted EventStatus = "Selesai"
)

type Event struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	Name        string      `gorm:"unique;not null" json:"name"`
	Description string      `gorm:"type:text" json:"description"`
	Location    string      `gorm:"type:varchar(255)" json:"location"`
	Date        Date        `gorm:"not null" json:"date"`
	Capacity    int         `gorm:"not null;check:capacity >= 0" json:"capacity"`
	Price       float64     `gorm:"not null;check:price >= 0" json:"price"`
	Status      EventStatus `gorm:"default:Aktif" json:"status"`

	Tickets []Ticket `gorm:"foreignKey:EventID" json:"tickets,omitempty"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"created_at"` // Auto-set saat create
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updated_at"` // Auto-update saat save/update
}
