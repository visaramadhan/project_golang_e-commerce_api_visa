package cart

import (
	"errors"

	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/product"
)

type CartService interface {
	AddProductToCart(userID, productID string, quantity int) error
	ListCart(userID string) ([]CartProductResponse, error)
	UpdateCart(userID, cartID string, quantity int) error
	DeleteCart(userID, cartID string) error
}

type cartService struct {
	repo           CartRepository
	productService product.ProductService
}

func NewCartService(repo CartRepository, productService product.ProductService) CartService {
	return &cartService{
		repo:           repo,
		productService: productService,
	}
}

func (s *cartService) AddProductToCart(userID, productID string, quantity int) error {
	if quantity <= 0 {
		quantity = 1
	}

	cart := Cart{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}

	return s.repo.AddToCart(cart)
}

func (s *cartService) ListCart(userID string) ([]CartProductResponse, error) {
	carts, err := s.repo.ListCart(userID)
	if err != nil {
		return nil, err
	}

	var cartResponse []CartProductResponse
	for _, cart := range carts {
		product, err := s.productService.GetProductByID(cart.ProductID)
		if err != nil {
			return nil, err
		}

		cartResponse = append(cartResponse, CartProductResponse{
			ID:         cart.ID,
			Name:       product.Name,
			Price:      product.Price,
			Quantity:   cart.Quantity,
			FinalPrice: product.Price * float64(cart.Quantity),
			Image:      product.Image,
		})
	}
	return cartResponse, nil
}

func (s *cartService) UpdateCart(userID, cartID string, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	cart := Cart{
		ID:       cartID,
		UserID:   userID,
		Quantity: quantity,
	}

	return s.repo.UpdateCart(cart)
}

func (s *cartService) DeleteCart(userID, cartID string) error {
	return s.repo.DeleteCart(cartID, userID)
}
