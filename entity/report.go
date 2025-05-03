package entity

// Gunakan gorm.Model jika perlu disimpan di database
type EventTicketReport struct {
	EventID      uint    `json:"event_id"`
	EventName    string  `json:"event_name"`
	TicketsSold  int     `json:"tickets_sold"`
	TotalRevenue float64 `json:"total_revenue"`
}

// RevenueSummary hanya untuk response JSON, tidak untuk database
type RevenueSummary struct {
	Period       string              `json:"period"`
	StartDate    string              `json:"start_date,omitempty"`
	EndDate      string              `json:"end_date,omitempty"`
	TotalRevenue float64             `json:"total_revenue"`
	TotalTickets int                 `json:"total_tickets"`
	Events       []EventTicketReport `json:"events,omitempty" gorm:"-"` // tanda '-' artinya diabaikan oleh GORM
}
