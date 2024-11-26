package product

import (
	"fmt"
	"log"
	"time"

	"github.com/visaramadhan/project_golang_e-commerce_api_visa/dto"
)

type ProductService interface {
	// GetAllProducts(limit, offset int) ([]ProductResponse, error)
	GetAllProducts(limit, page int) (PaginatedProductResponse, error)
	GetBestSellingProducts(month, year, page, pageSize int) ([]Product, error)
	// UpdatePromoProduct(productID string, dto dto.PromoProductDTO) error
	UpdatePromoProduct(productID string, dto dto.PromoProductDTO) (*Product, error)
	// Save(product *Product) (*Product, error)
	GetRecommendedProducts() ([]dto.RecommendProductDTO, error)
	GetProducts(name string, page, limit int) ([]Product, int, error)
	GetProductsByCategory(categoryID string, limit, page int) ([]Product, error)
	GetProductByID(productID string) (Product, error)
	GetProductsByID(productID string) (ProductDetail, error)
}

type productService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &productService{repo: repo}
}

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Discount    float64 `json:"discount"`
	FinalPrice  float64 `json:"final_price"`
	Image       string  `json:"image"`
	IsNew       bool    `json:"is_new"`
	Rating      float64 `json:"rating"`
	TotalRating int     `json:"total_rating"`
}

type PaginatedProductResponse struct {
	Products    []ProductResponse `json:"products"`
	TotalData   int64             `json:"total_data"`
	TotalPages  int               `json:"total_pages"`
	CurrentPage int               `json:"current_page"`
}

func (s *productService) GetProductsByID(productID string) (ProductDetail, error) {
	return s.repo.GetProductsByID(productID)
}

func (s *productService) GetAllProducts(limit, page int) (PaginatedProductResponse, error) {
	offset := (page - 1) * limit

	totalData, err := s.repo.GetProductCount()
	if err != nil {
		return PaginatedProductResponse{}, err
	}

	products, err := s.repo.GetAllProducts(limit, offset)
	if err != nil {
		return PaginatedProductResponse{}, err
	}

	var response []ProductResponse
	for _, p := range products {
		isNew := time.Since(p.CreatedAt).Hours() < 720
		finalPrice := p.Price * (1 - (p.Discount / 100))
		response = append(response, ProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Discount:    p.Discount,
			FinalPrice:  finalPrice,
			Image:       p.Image,
			IsNew:       isNew,
			Rating:      p.Rating,
			TotalRating: p.TotalRating,
		})
	}

	totalPages := int((totalData + int64(limit) - 1) / int64(limit))

	return PaginatedProductResponse{
		Products:    response,
		TotalData:   totalData,
		TotalPages:  totalPages,
		CurrentPage: page,
	}, nil
}

func (s *productService) GetProductByID(productID string) (Product, error) {
	return s.repo.GetProductByID(productID)
}

func (s *productService) GetBestSellingProducts(month, year, page, pageSize int) ([]Product, error) {
	if month == 0 {
		month = int(time.Now().Month())
	}
	if year == 0 {
		year = time.Now().Year()
	}

	products, err := s.repo.GetBestSellingProducts(month, year, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get best selling products: %v", err)
	}

	return products, nil
}

func (s *productService) UpdatePromoProduct(productID string, dto dto.PromoProductDTO) (*Product, error) {
	product, err := s.repo.GetProductByID(productID)
	if err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	product.DiscountStartDate = dto.DiscountStartDate
	product.DiscountEndDate = dto.DiscountEndDate
	product.DiscountPercentage = dto.DiscountPercentage
	product.IsPromo = dto.IsPromo

	updatedProduct, err := s.repo.Save(&product)
	if err != nil {
		return nil, fmt.Errorf("failed to save updated product: %w", err)
	}

	return updatedProduct, nil
}

func (s *productService) GetRecommendedProducts() ([]dto.RecommendProductDTO, error) {
	products, err := s.repo.GetRecommendedProducts()
	if err != nil {
		return nil, err
	}

	var recommendedProducts []dto.RecommendProductDTO
	for _, product := range products {
		recommendedProducts = append(recommendedProducts, dto.RecommendProductDTO{
			Title:     product.Name,
			Subtitle:  product.Description,
			Photo:     product.Image,
			ProductID: product.ID,
		})
	}
	fmt.Printf("Recommended Products: %+v\n", recommendedProducts)

	return recommendedProducts, nil
}

func (s *productService) GetProducts(name string, page, limit int) ([]Product, int, error) {
	offset := (page - 1) * limit
	products, total, err := s.repo.FindAll(name, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	totalPages := (int(total) + limit - 1) / limit
	return products, totalPages, nil
}

func (s *productService) GetProductsByCategory(categoryID string, limit, page int) ([]Product, error) {
	offset := (page - 1) * limit
	products, err := s.repo.GetProductsByCategory(categoryID, limit, offset)
	log.Printf("categoryID: %s, limit: %d, offset: %d", categoryID, limit, offset)

	if err != nil {
		return nil, err
	}
	return products, nil
}
