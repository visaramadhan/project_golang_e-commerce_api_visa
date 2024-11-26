package wishlist

import 	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/product"


type WishlistService interface {
	AddProductToWishlist(userID string, productID string) error
	GetAllWishlist(userID string) ([]product.ProductResponse, error)
	DeleteWishlist(userID, productID string) error
}

type wishlistService struct {
	repo WishlistRepository
}

func NewWishlistService(repo WishlistRepository) WishlistService {
	return &wishlistService{repo: repo}
}

func (s *wishlistService) AddProductToWishlist(userID string, productID string) error {
	return s.repo.AddToWishlist(userID, productID)
}

func (s *wishlistService) GetAllWishlist(userID string) ([]product.ProductResponse, error) {
	return s.repo.GetAllWishlist(userID)
}

func (s *wishlistService) DeleteWishlist(userID, productID string) error {
	return s.repo.DeleteWishlist(userID, productID)
}
