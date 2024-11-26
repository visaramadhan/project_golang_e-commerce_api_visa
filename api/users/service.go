package users

import (
	"errors"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/helper"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterInput) (*Users, error)
	Login(input LoginInput) (map[string]interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

type RegisterInput struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email,omitempty" binding:"omitempty,email"`
	PhoneNumber string `json:"phoneNumber,omitempty" binding:"omitempty,numeric"`
	Password    string `json:"password" binding:"required,min=6"`
}

func (s *service) Register(input RegisterInput) (*Users, error) {
	if input.Email == "" && input.PhoneNumber == "" {
		return nil, errors.New("email or phone number must be provided")
	}

	if input.Email != "" {
		existingUser, _ := s.repository.FindByEmail(input.Email)
		if existingUser != nil {
			return nil, errors.New("email already registered")
		}
	}

	if input.PhoneNumber != "" {
		existingUser, _ := s.repository.FindByPhoneNumber(input.PhoneNumber)
		if existingUser != nil {
			return nil, errors.New("phone number already registered")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &Users{
		ID:          uuid.New().String(),
		Name:        input.Name,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    string(hashedPassword),
	}

	if err := s.repository.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

type LoginInput struct {
	EmailOrPhone string `json:"emailOrPhone" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

func (s *service) Login(input LoginInput) (map[string]interface{}, error) {
	user, err := s.repository.FindByEmailOrPhone(input.EmailOrPhone)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("email or phone number not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := helper.GenerateToken(user.ID)//helper(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	response := map[string]interface{}{
		"token": token,
		"session": map[string]interface{}{
			"user_id":      user.ID,
			"name":         user.Name,
			"email":        user.Email,
			"phone_number": user.PhoneNumber,
		},
	}

	return response, nil
}
