package users

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Mock gorm.DB
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	result := &gorm.DB{}
	if err, ok := args.Get(0).(error); ok {
		result.Error = err
	}
	return result
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	m.Called(query, args)
	return &gorm.DB{}
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	result := &gorm.DB{}
	if err, ok := args.Get(0).(error); ok {
		result.Error = err
	}
	return result
}

func TestSave(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewRepository(mockDB.Create(mockDB))
	user := &Users{Email: "jefri@mail.com.com", PhoneNumber: "123456789"}
	mockDB.On("Create", user).Return(nil)
	err := repo.Save(user)
	assert.Nil(t, err)

	mockDB.On("Create", user).Return(errors.New("duplicate key value violates unique constraint"))
	err = repo.Save(user)
	assert.EqualError(t, err, "email or phone number already registered")
}

// func TestFindByEmail(t *testing.T) {
// 	mockDB := new(MockDB) // Menggunakan MockDB, bukan MockRepository
// 	repo := &repository{db: mockDB}

// 	expectedUser := &Users{Email: "test@example.com"}
// 	mockDB.On("First", mock.Anything, []interface{}{"test@example.com"}).Return(nil).Run(func(args mock.Arguments) {
// 		arg := args.Get(0).(*Users)
// 		*arg = *expectedUser
// 	})
// 	user, err := repo.FindByEmail("test@example.com")
// 	assert.Nil(t, err)
// 	assert.Equal(t, expectedUser, user)

// 	mockDB.On("First", mock.Anything, []interface{}{"notfound@example.com"}).Return(gorm.ErrRecordNotFound)
// 	user, err = repo.FindByEmail("notfound@example.com")
// 	assert.Nil(t, err)
// 	assert.Nil(t, user)
// }

// func TestFindByPhoneNumber(t *testing.T) {
// 	mockDB := new(MockDB) // Menggunakan MockDB, bukan MockRepository
// 	repo := &repository{db: mockDB}

// 	expectedUser := &Users{PhoneNumber: "123456789"}
// 	mockDB.On("First", mock.Anything, []interface{}{"123456789"}).Return(nil).Run(func(args mock.Arguments) {
// 		arg := args.Get(0).(*Users)
// 		*arg = *expectedUser
// 	})
// 	user, err := repo.FindByPhoneNumber("123456789")
// 	assert.Nil(t, err)
// 	assert.Equal(t, expectedUser, user)

// 	mockDB.On("First", mock.Anything, []interface{}{"987654321"}).Return(gorm.ErrRecordNotFound)
// 	user, err = repo.FindByPhoneNumber("987654321")
// 	assert.Nil(t, err)
// 	assert.Nil(t, user)
// }

// func TestFindByEmailOrPhone(t *testing.T) {
// 	mockDB := new(MockDB) // Menggunakan MockDB, bukan MockRepository
// 	repo := &repository{db: mockDB}

// 	expectedUser := &Users{Email: "test@example.com"}
// 	mockDB.On("First", mock.Anything, []interface{}{"test@example.com", "test@example.com"}).Return(nil).Run(func(args mock.Arguments) {
// 		arg := args.Get(0).(*Users)
// 		*arg = *expectedUser
// 	})
// 	user, err := repo.FindByEmailOrPhone("test@example.com")
// 	assert.Nil(t, err)
// 	assert.Equal(t, expectedUser, user)

// 	mockDB.On("First", mock.Anything, []interface{}{"notfound@example.com", "notfound@example.com"}).Return(gorm.ErrRecordNotFound)
// 	user, err = repo.FindByEmailOrPhone("notfound@example.com")
// 	assert.Nil(t, err)
// 	assert.Nil(t, user)
// }
