package service

import (
	"errors"
	"ticket-system/entity"
	"ticket-system/repository"
	"time"
)

type ReportService interface {
	GenerateEventReport(eventID uint) (*entity.EventTicketReport, error)
	GenerateRevenueSummary(period string, startDate, endDate string) (*entity.RevenueSummary, error)
}

type reportService struct {
	repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) ReportService {
	return &reportService{repo: repo}
}

func (s *reportService) GenerateEventReport(eventID uint) (*entity.EventTicketReport, error) {
	return s.repo.GetEventTicketReport(eventID)
}

func (s *reportService) GenerateRevenueSummary(period string, startDate, endDate string) (*entity.RevenueSummary, error) {
	// Validate period
	validPeriods := map[string]bool{
		"day": true, "week": true, "month": true, "year": true, "custom": true,
	}
	if !validPeriods[period] {
		return nil, errors.New("invalid period parameter")
	}

	var start, end time.Time
	var err error

	if period == "custom" {
		start, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			return nil, errors.New("invalid start_date format, use YYYY-MM-DD")
		}

		end, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			return nil, errors.New("invalid end_date format, use YYYY-MM-DD")
		}
	} else {
		now := time.Now()
		switch period {
		case "day":
			start = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			end = start.Add(24 * time.Hour)
		case "week":
			start = now.AddDate(0, 0, -int(now.Weekday()))
			end = start.AddDate(0, 0, 7)
		case "month":
			start = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
			end = start.AddDate(0, 1, 0)
		case "year":
			start = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
			end = start.AddDate(1, 0, 0)
		}
	}

	return s.repo.GetRevenueSummary(period, start, end)
}
