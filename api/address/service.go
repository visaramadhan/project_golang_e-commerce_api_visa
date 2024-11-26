package address

type AddressService interface {
	GetAddresses(userID string) ([]Address, error)
	CreateAddress(userID string, address Address) (Address, error)
	UpdateAddress(userID, addressID string, address Address) (Address, error)
	DeleteAddress(userID, addressID string) error
}

type addressService struct {
	repo AddressRepository
}

func NewAddressService(repo AddressRepository) AddressService {
	return &addressService{repo}
}

func (s *addressService) GetAddresses(userID string) ([]Address, error) {
	return s.repo.GetAddressesByUserID(userID)
}

func (s *addressService) CreateAddress(userID string, address Address) (Address, error) {
	return s.repo.CreateAddress(userID, address)
}

func (s *addressService) UpdateAddress(userID, addressID string, address Address) (Address, error) {
	return s.repo.UpdateAddress(userID, addressID, address)
}

func (s *addressService) DeleteAddress(userID, addressID string) error {
	return s.repo.DeleteAddress(userID, addressID)
}
