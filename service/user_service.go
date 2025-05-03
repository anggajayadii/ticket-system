package service

import (
	"errors"
	"ticket-system/dto"
	"ticket-system/entity"
	"ticket-system/repository"
	password "ticket-system/utils"
)

var (
	ErrEmailExists  = errors.New("email already registered")
	ErrInvalidCreds = errors.New("invalid email or password")
)

type AuthService interface {
	Register(input dto.RegisterRequest) (*entity.User, error)
	Login(input dto.LoginRequest) (*entity.User, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(input dto.RegisterRequest) (*entity.User, error) {
	// Cek email sudah ada
	exists, err := s.repo.EmailExists(input.Email)
	if err != nil || exists {
		return nil, ErrEmailExists
	}

	// Hash password
	hashed, err := password.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
		Role:     input.Role,
	}

	return s.repo.Create(user)
}

func (s *authService) Login(input dto.LoginRequest) (*entity.User, error) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil || user == nil {
		return nil, ErrInvalidCreds
	}

	if err := password.Compare(user.Password, input.Password); err != nil {
		return nil, ErrInvalidCreds
	}

	return user, nil
}
