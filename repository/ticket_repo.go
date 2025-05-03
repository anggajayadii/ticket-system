package repository

import (
	"errors"
	"ticket-system/entity"

	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ticket *entity.Ticket) error
	FindByID(id uint) (*entity.Ticket, error)
	FindByUserID(userID uint) ([]entity.Ticket, error)
	FindAll() ([]entity.Ticket, error)
	Update(ticket *entity.Ticket) error
	Delete(ticket *entity.Ticket) error
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db: db}
}

func (r *ticketRepository) Create(ticket *entity.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) FindByID(id uint) (*entity.Ticket, error) {
	if id == 0 {
		return nil, errors.New("ID cannot be zero")
	}

	var ticket entity.Ticket
	err := r.db.Preload("User").Preload("Event").First(&ticket, id).Error
	if err != nil {
		return nil, err // Selalu return error, tidak return nil
	}
	return &ticket, nil
}

// func (r *ticketRepo) FindByID(id uint) (*entity.Ticket, error) {
//     var ticket entity.Ticket
//     err := r.db.Where("id = ?", id).First(&ticket).Error

func (r *ticketRepository) FindByUserID(userID uint) ([]entity.Ticket, error) {
	var tickets []entity.Ticket
	err := r.db.Preload("User").Preload("Event").Where("user_id = ?", userID).Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *ticketRepository) FindAll() ([]entity.Ticket, error) {
	var tickets []entity.Ticket
	err := r.db.Preload("User").Preload("Event").Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *ticketRepository) Update(ticket *entity.Ticket) error {
	return r.db.Save(ticket).Error
}

func (r *ticketRepository) Delete(ticket *entity.Ticket) error {
	return r.db.Delete(ticket).Error
}
