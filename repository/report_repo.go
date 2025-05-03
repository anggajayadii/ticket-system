package repository

import (
	"ticket-system/entity"
	"time"

	"gorm.io/gorm"
)

type ReportRepository interface {
	GetEventTicketReport(eventID uint) (*entity.EventTicketReport, error)
	GetRevenueSummary(period string, startDate, endDate time.Time) (*entity.RevenueSummary, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) GetEventTicketReport(eventID uint) (*entity.EventTicketReport, error) {
	var report entity.EventTicketReport

	err := r.db.Table("tickets").
		Select("events.id as event_id, events.name as event_name, sum(tickets.quantity) as tickets_sold, sum(tickets.total) as total_revenue").
		Joins("left join events on events.id = tickets.event_id").
		Where("tickets.event_id = ? AND tickets.status != ?", eventID, "dibatalkan").
		Group("events.id").
		Scan(&report).Error

	return &report, err
}

func (r *reportRepository) GetRevenueSummary(period string, startDate, endDate time.Time) (*entity.RevenueSummary, error) {
	var summary entity.RevenueSummary
	summary.Period = period

	query := r.db.Table("tickets").
		Select("sum(total) as total_revenue, sum(quantity) as total_tickets").
		Where("status != ?", "dibatalkan")

	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
		summary.StartDate = startDate.Format("2006-01-02")
		summary.EndDate = endDate.Format("2006-01-02")
	}

	err := query.Scan(&summary).Error
	if err != nil {
		return nil, err
	}

	// Get detail per event
	var eventReports []entity.EventTicketReport
	eventQuery := r.db.Table("tickets").
		Select("events.id as event_id, events.name as event_name, sum(tickets.quantity) as tickets_sold, sum(tickets.total) as total_revenue").
		Joins("left join events on events.id = tickets.event_id").
		Where("tickets.status != ?", "dibatalkan")

	if !startDate.IsZero() && !endDate.IsZero() {
		eventQuery = eventQuery.Where("tickets.created_at BETWEEN ? AND ?", startDate, endDate)
	}

	err = eventQuery.Group("events.id").Scan(&eventReports).Error
	if err != nil {
		return nil, err
	}

	summary.Events = eventReports

	return &summary, nil
}
