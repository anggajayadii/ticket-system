package service

import (
	"ticket-system/entity"
	"ticket-system/repository"
	"time"
)

type EventService interface {
	CreateEvent(event *entity.Event) (*entity.Event, error)
	GetAllEvents() ([]entity.Event, error)
	GetEventByID(id uint) (*entity.Event, error)
	UpdateEvent(event *entity.Event) (*entity.Event, error)
	DeleteEvent(id uint) error
}

type eventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) EventService {
	return &eventService{repo}
}

func (s *eventService) CreateEvent(event *entity.Event) (*entity.Event, error) {
	// Set default status if not provided
	if event.Status == "" {
		event.Status = entity.StatusActive
	}

	err := s.repo.Create(event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (s *eventService) GetAllEvents() ([]entity.Event, error) {
	return s.repo.FindAll()
}

func (s *eventService) GetEventByID(id uint) (*entity.Event, error) {
	return s.repo.FindByID(id)
}

func (s *eventService) UpdateEvent(event *entity.Event) (*entity.Event, error) {
	existingEvent, err := s.repo.FindByID(event.ID)
	if err != nil {
		return nil, err
	}

	// Apply partial updates
	if event.Name != "" {
		existingEvent.Name = event.Name
	}
	if event.Description != "" {
		existingEvent.Description = event.Description
	}
	if event.Location != "" {
		existingEvent.Location = event.Location
	}
	// Periksa apakah Date tidak zero
	if !time.Time(event.Date).IsZero() {
		existingEvent.Date = event.Date
	}
	if event.Capacity != 0 {
		existingEvent.Capacity = event.Capacity
	}
	if event.Price != 0 {
		existingEvent.Price = event.Price
	}
	if event.Status != "" {
		existingEvent.Status = event.Status
	}

	err = s.repo.Update(existingEvent)
	if err != nil {
		return nil, err
	}

	return existingEvent, nil
}

func (s *eventService) DeleteEvent(id uint) error {
	return s.repo.Delete(id)
}
