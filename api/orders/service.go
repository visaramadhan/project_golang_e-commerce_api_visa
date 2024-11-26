package orders

type OrderService interface {
	CreateOrder(userID string, order Orderes) (OrderResponse, error)
}

type orderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &orderService{repo}
}

func (s *orderService) CreateOrder(userID string, order Orderes) (OrderResponse, error) {
	return s.repo.CreateOrder(userID, order)
}
