package entity

import "time"

type Ticket struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	UserID   uint    `gorm:"not null" json:"user_id"`
	EventID  uint    `gorm:"not null" json:"event_id"`
	Status   string  `gorm:"type:enum('tersedia','habis','dibatalkan');default:'tersedia'" json:"status"`
	Quantity int     `gorm:"not null;check:quantity > 0" json:"quantity"`
	Total    float64 `gorm:"not null" json:"total"`

	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Event Event `gorm:"foreignKey:EventID" json:"event,omitempty"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"created_at"` // Auto-set saat create
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updated_at"` // Auto-update saat save/update

}
