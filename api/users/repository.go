package users

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user *Users) error
	FindByEmail(email string) (*Users, error)
	FindByPhoneNumber(phone string) (*Users, error)
	FindByEmailOrPhone(emailOrPhone string) (*Users, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// func (r *repository) Save(user *Users) error {
// 	return r.db.Create(user).Error
// }

func (r *repository) Save(user *Users) error {
	err := r.db.Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errors.New("email or phone number already registered")
		}
		return err
	}
	return nil
}

func (r *repository) FindByEmail(email string) (*Users, error) {
	var user Users
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *repository) FindByPhoneNumber(phone string) (*Users, error) {
	var user Users
	err := r.db.Where("phone_number = ?", phone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *repository) FindByEmailOrPhone(emailOrPhone string) (*Users, error) {
	var user Users
	err := r.db.Where("email = ? OR phone_number = ?", emailOrPhone, emailOrPhone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}
