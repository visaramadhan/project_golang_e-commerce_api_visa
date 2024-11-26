package wishlist

import (
	"time"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/product"

)

type Wishlist struct {
	ID        string          `gorm:"primaryKey"`
	UserID    string          `json:"user_id"`
	ProductID string          `json:"product_id"`
	Product   product.Product `gorm:"foreignKey:ProductID"`
	CreatedAt time.Time       `json:"created_at"`
}
