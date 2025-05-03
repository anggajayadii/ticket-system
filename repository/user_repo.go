package repository

import (
	"errors"
	"ticket-system/entity"

	"gorm.io/gorm"
)

var (
	ErrDuplicateEmail = errors.New("email already exists")
)

type AuthRepository interface {
	EmailExists(email string) (bool, error)
	Create(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

// EmailExists mengecek apakah email sudah terdaftar
func (r *authRepository) EmailExists(email string) (bool, error) {
	var count int64
	err := r.db.Model(&entity.User{}).
		Where("email = ?", email).
		Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Create menyimpan user baru ke database
func (r *authRepository) Create(user *entity.User) (*entity.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		if isDuplicateError(err) {
			return nil, ErrDuplicateEmail
		}
		return nil, err
	}
	return user, nil
}

// FindByEmail mencari user berdasarkan email
func (r *authRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// isDuplicateError mengecek apakah error berasal dari duplikasi unique key
func isDuplicateError(err error) bool {
	// Implementasi bisa disesuaikan dengan driver database
	// Contoh untuk MySQL:
	return errors.Is(err, gorm.ErrDuplicatedKey)
}
