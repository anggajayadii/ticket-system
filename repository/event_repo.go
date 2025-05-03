package repository

import (
	"ticket-system/entity"
	"time"

	"gorm.io/gorm"
)

type EventRepository interface {
	Create(event *entity.Event) error
	FindAll() ([]entity.Event, error)
	FindByID(id uint) (*entity.Event, error)
	Update(event *entity.Event) error
	Delete(id uint) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) Create(event *entity.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) FindAll() ([]entity.Event, error) {
	var events []entity.Event
	err := r.db.Find(&events).Error
	return events, err
}

func (r *eventRepository) FindByID(id uint) (*entity.Event, error) {
	var event entity.Event
	err := r.db.First(&event, id).Error
	return &event, err
}

func (r *eventRepository) Update(event *entity.Event) error {
	updateData := make(map[string]interface{})

	if event.Name != "" {
		updateData["name"] = event.Name
	}
	if event.Description != "" {
		updateData["description"] = event.Description
	}
	if event.Location != "" {
		updateData["location"] = event.Location
	}
	// Periksa apakah Date tidak zero
	if !time.Time(event.Date).IsZero() {
		updateData["date"] = event.Date
	}
	if event.Capacity != 0 {
		updateData["capacity"] = event.Capacity
	}
	if event.Price != 0 {
		updateData["price"] = event.Price
	}
	if event.Status != "" {
		updateData["status"] = event.Status
	}

	return r.db.Model(&entity.Event{}).Where("id = ?", event.ID).Updates(updateData).Error
}

func (r *eventRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Event{}, id).Error
}
