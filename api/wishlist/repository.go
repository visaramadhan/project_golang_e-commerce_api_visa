package wishlist

import (
	"time"

	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/product"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WishlistRepository interface {
	AddToWishlist(userID string, productID string) error
	GetAllWishlist(userID string) ([]product.ProductResponse, error)
	DeleteWishlist(userID, productID string) error
}

type wishlistRepository struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) WishlistRepository {
	return &wishlistRepository{db: db}
}

func (r *wishlistRepository) AddToWishlist(userID string, productID string) error {
	wishlistItem := Wishlist{
		ID:        uuid.NewString(),
		UserID:    userID,
		ProductID: productID,
		CreatedAt: time.Now(),
	}

	return r.db.Create(&wishlistItem).Error
}

func (r *wishlistRepository) GetAllWishlist(userID string) ([]product.ProductResponse, error) {
	var wishlists []product.ProductResponse
	err := r.db.Raw(`
		SELECT 
			w.product_id, 
			p.name, p.description, p.price, p.discount, 
			(p.price - (p.price * p.discount / 100)) AS final_price, 
			p.image, p.is_new, p.rating, p.total_rating
		FROM wishlists w
		JOIN products p ON w.product_id = p.id
		WHERE w.user_id = ?
	`, userID).Scan(&wishlists).Error
	return wishlists, err
}

func (r *wishlistRepository) DeleteWishlist(userID, productID string) error {
	return r.db.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&Wishlist{}).Error
}
