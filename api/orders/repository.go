package orders

import "gorm.io/gorm"

type OrderRepository interface {
	CreateOrder(userID string, order Orderes) (OrderResponse, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateOrder(userID string, order Orderes) (OrderResponse, error) {
	var orderResponse OrderResponse
	if err := r.db.Create(&order).Error; err != nil {
		return orderResponse, err
	}
	orderResponse.OrderID = order.ID
	orderResponse.Status = "pending"

	totalAmount := 0.0
	for _, item := range order.OrderItems {
		totalAmount += item.Price * float64(item.Quantity)
	}
	orderResponse.TotalAmount = totalAmount
	orderResponse.PaymentMethod = order.PaymentMethod
	return orderResponse, nil
}
