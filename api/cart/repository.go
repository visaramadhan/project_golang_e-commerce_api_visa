package cart

import (
	"time"

	"gorm.io/gorm"
)

type CartRepository interface {
	AddToCart(cart Cart) error
	ListCart(userID string) ([]Cart, error)
	UpdateCart(cart Cart) error
	DeleteCart(cartID, userID string) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) AddToCart(cart Cart) error {
	return r.db.Create(&cart).Error
}

func (r *cartRepository) ListCart(userID string) ([]Cart, error) {
	var carts []Cart
	if err := r.db.Where("user_id = ?", userID).Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *cartRepository) UpdateCart(cart Cart) error {
	return r.db.Model(&Cart{}).
		Where("id = ? AND user_id = ?", cart.ID, cart.UserID).
		Updates(map[string]interface{}{
			"quantity":   cart.Quantity,
			"updated_at": time.Now(),
		}).Error
}

func (r *cartRepository) DeleteCart(cartID, userID string) error {
	return r.db.Where("id = ? AND user_id = ?", cartID, userID).Delete(&Cart{}).Error
}
