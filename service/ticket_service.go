package service

import (
	"errors"
	"ticket-system/entity"
	"ticket-system/repository"
	"time"
)

type TicketService interface {
	PurchaseTicket(ticket *entity.Ticket) (*entity.Ticket, error)
	GetMyTickets(userID uint) ([]entity.Ticket, error)
	GetAllTickets() ([]entity.Ticket, error)
	GetTicketsByUserID(userID uint) ([]entity.Ticket, error)
	CancelTicket(ticketID uint, userID uint, isAdmin bool) error
}

type ticketService struct {
	ticketRepo repository.TicketRepository
	eventRepo  repository.EventRepository
}

func NewTicketService(ticketRepo repository.TicketRepository, eventRepo repository.EventRepository) TicketService {
	return &ticketService{
		ticketRepo: ticketRepo,
		eventRepo:  eventRepo,
	}
}

func (s *ticketService) PurchaseTicket(ticket *entity.Ticket) (*entity.Ticket, error) {
	// Validate ticket data
	if ticket.Quantity <= 0 {
		return nil, errors.New("quantity must be greater than 0")
	}

	// Check if event exists and has available tickets
	event, err := s.eventRepo.FindByID(ticket.EventID)
	if err != nil {
		return nil, err
	}
	if event == nil {
		return nil, errors.New("event not found")
	}
	// Hitung total harga
	ticket.Total = event.Price * float64(ticket.Quantity)

	// Set default status and timestamps
	ticket.Status = "tersedia"
	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()

	// Create the ticket
	err = s.ticketRepo.Create(ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (s *ticketService) GetMyTickets(userID uint) ([]entity.Ticket, error) {
	return s.ticketRepo.FindByUserID(userID)
}

func (s *ticketService) GetAllTickets() ([]entity.Ticket, error) {
	return s.ticketRepo.FindAll()
}

func (s *ticketService) GetTicketsByUserID(userID uint) ([]entity.Ticket, error) {
	return s.ticketRepo.FindByUserID(userID)
}

func (s *ticketService) CancelTicket(ticketID uint, userID uint, isAdmin bool) error {
	ticket, err := s.ticketRepo.FindByID(ticketID)
	if err != nil {
		return err
	}
	if ticket == nil {
		return errors.New("ticket not found")
	}

	// Check if user is the owner or admin
	if !isAdmin && ticket.UserID != userID {
		return errors.New("unauthorized to cancel this ticket")
	}

	// Update ticket status
	ticket.Status = "dibatalkan"
	ticket.UpdatedAt = time.Now()

	return s.ticketRepo.Update(ticket)
}
