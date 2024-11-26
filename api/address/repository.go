package address

import "gorm.io/gorm"

type AddressRepository interface {
	GetAddressesByUserID(userID string) ([]Address, error)
	CreateAddress(userID string, address Address) (Address, error)
	UpdateAddress(userID, addressID string, address Address) (Address, error)
	DeleteAddress(userID, addressID string) error
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db}
}

func (r *addressRepository) GetAddressesByUserID(userID string) ([]Address, error) {
	var addresses []Address
	if err := r.db.Where("user_id = ?", userID).Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}

func (r *addressRepository) CreateAddress(userID string, address Address) (Address, error) {
	address.UserID = userID
	if err := r.db.Create(&address).Error; err != nil {
		return Address{}, err
	}
	return address, nil
}

func (r *addressRepository) UpdateAddress(userID, addressID string, address Address) (Address, error) {
	var existing Address
	if err := r.db.Where("user_id = ? AND id = ?", userID, addressID).First(&existing).Error; err != nil {
		return Address{}, err
	}
	if err := r.db.Model(&existing).Updates(address).Error; err != nil {
		return Address{}, err
	}
	return existing, nil
}

func (r *addressRepository) DeleteAddress(userID, addressID string) error {
	var address Address
	if err := r.db.Where("user_id = ? AND id = ?", userID, addressID).First(&address).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&address).Error; err != nil {
		return err
	}
	return nil
}
