package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(user *Users) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockRepository) FindByEmail(email string) (*Users, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Users), args.Error(1)
}

func (m *MockRepository) FindByPhoneNumber(phone string) (*Users, error) {
	args := m.Called(phone)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Users), args.Error(1)
}

func (m *MockRepository) FindByEmailOrPhone(emailOrPhone string) (*Users, error) {
	args := m.Called(emailOrPhone)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Users), args.Error(1)
}

func TestRegister(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		input         RegisterInput
		mockBehavior  func()
		expectedError string
	}{
		{
			name: "success register with email",
			input: RegisterInput{
				Name:        "Jefri",
				Email:       "jefri@mail.com",
				PhoneNumber: "",
				Password:    "password123",
			},
			mockBehavior: func() {
				mockRepo.On("FindByEmail", "jefri@mail.com").Return(nil, nil)
				mockRepo.On("Save", mock.Anything).Return(nil)
			},
			expectedError: "",
		},
		{
			name: "email already registered",
			input: RegisterInput{
				Name:        "Jefri",
				Email:       "jefri@mail.com",
				PhoneNumber: "",
				Password:    "password123",
			},
			mockBehavior: func() {
				mockRepo.On("FindByEmail", "jefri@mail.com").Return(&Users{}, nil)
			},
			expectedError: "email already registered",
		},
		{
			name: "phone number already registered",
			input: RegisterInput{
				Name:        "Jefri",
				Email:       "",
				PhoneNumber: "123456789",
				Password:    "password123",
			},
			mockBehavior: func() {
				mockRepo.On("FindByPhoneNumber", "123456789").Return(&Users{}, nil)
			},
			expectedError: "phone number already registered",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.ExpectedCalls = nil

			tt.mockBehavior()

			user, err := service.Register(tt.input)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestLogin(t *testing.T) {
	// mocktime.Set(time.Unix(1732550911, 0))
	// defer mocktime.Reset()
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		input         LoginInput
		mockBehavior  func()
		expectedError string
		expectedToken string
	}{
		{
			name: "success login",
			input: LoginInput{
				EmailOrPhone: "jefri@mail.com",
				Password:     "password123",
			},
			mockBehavior: func() {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
				mockRepo.On("FindByEmailOrPhone", "jefri@mail.com").Return(&Users{
					Password: string(hashedPassword),
				}, nil)
			},
			expectedError: "",
			expectedToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzI1NTA5MTEsInVzZXJfaWQiOiI1NDMyOWRmMi05Mjg3LTRlMDQtYjFlOS1kYmY2MDk2NmYzMjgifQ.GTOKl58co0-HIdk1CRSyfCj1SC2LQf8wBMPKVn3Q9Ws",
		},
		{
			name: "user not found",
			input: LoginInput{
				EmailOrPhone: "unknown@example.com",
				Password:     "password123",
			},
			mockBehavior: func() {
				mockRepo.On("FindByEmailOrPhone", "unknown@example.com").Return(nil, nil)
			},
			expectedError: "email or phone number not found",
		},
		{
			name: "invalid password",
			input: LoginInput{
				EmailOrPhone: "jefri@mail.com",
				Password:     "wrongpassword",
			},
			mockBehavior: func() {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
				mockRepo.On("FindByEmailOrPhone", "jefri@mail.com").Return(&Users{
					Password: string(hashedPassword),
				}, nil)
			},
			expectedError: "invalid password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.ExpectedCalls = nil

			tt.mockBehavior()

			response, err := service.Login(tt.input)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.NotEqual(t, tt.expectedToken, response["token"])
			}

			mockRepo.AssertExpectations(t)
		})
	}
}