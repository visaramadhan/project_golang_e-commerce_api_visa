package orders

type Orderes struct {
	ID            string      `json:"id" gorm:"primaryKey"`
	UserID        string      `json:"user_id" gorm:"index"`
	AddressID     string      `json:"address_id" gorm:"index"`
	PaymentMethod string      `json:"payment_method"`
	OrderItems    []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        string  `json:"id" gorm:"primaryKey"`
	OrderID   string  `json:"order_id" gorm:"index"`
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderResponse struct {
	OrderID       string  `json:"order_id"`
	Status        string  `json:"status"`
	TotalAmount   float64 `json:"total_amount"`
	PaymentMethod string  `json:"payment_method"`
}
