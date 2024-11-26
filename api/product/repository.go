package product

import (
	"log"

	"github.com/visaramadhan/project_golang_e-commerce_api_visa/dto"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts(limit, offset int) ([]Product, error)
	GetProductCount() (int64, error)
	GetBestSellingProducts(month, year, page, pageSize int) ([]Product, error)
	GetProductByID(productID string) (Product, error)
	UpdatePromoProduct(productID string, dto dto.PromoProductDTO) error
	Save(product *Product) (*Product, error)
	GetRecommendedProducts() ([]Product, error)
	FindAll(name string, limit, offset int) ([]Product, int64, error)
	GetProductsByCategory(categoryID string, limit, offset int) ([]Product, error)
	GetProductsByID(productID string) (ProductDetail, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAllProducts(limit, offset int) ([]Product, error) {
	var products []Product
	err := r.db.Order("RANDOM()").Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

func (r *productRepository) GetProductCount() (int64, error) {
	var count int64
	if err := r.db.Model(&Product{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetBestSellingProducts(db *gorm.DB, month int, year int, page int, pageSize int) ([]Product, error) {
	var products []Product
	err := db.Joins("JOIN orders ON products.id = orders.product_id").
		Where("EXTRACT(MONTH FROM orders.created_at) = ? AND EXTRACT(YEAR FROM orders.created_at) = ?", month, year).
		Group("products.id").
		Order("SUM(orders.quantity) DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&products).Error
	return products, err
}

func (r *productRepository) GetBestSellingProducts(month, year, page, pageSize int) ([]Product, error) {
	var products []Product
	err := r.db.Joins("JOIN orders ON products.id = orders.product_id").
		Where("EXTRACT(MONTH FROM orders.created_at) = ? AND EXTRACT(YEAR FROM orders.created_at) = ?", month, year).
		Group("products.id").
		Order("SUM(orders.quantity) DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&products).Error
	return products, err
}

func (r *productRepository) GetProductsByID(productID string) (ProductDetail, error) {
	var product ProductDetail
	query := `
SELECT pd.id, pd.name, pd.description, pd.rating, pd.total_rating,
       ARRAY_AGG(DISTINCT pi.image_url) AS images,
       json_agg(DISTINCT jsonb_build_object('id', v.id, 'color', v.color, 'size', v.size)) AS variations
FROM product_details pd
LEFT JOIN variations v ON v.product_id = pd.id
WHERE pd.id = ?
GROUP BY pd.id;
`
	err := r.db.Raw(query, productID).Scan(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) GetProductByID(productID string) (Product, error) {
	var product Product
	if err := r.db.Where("id = ?", productID).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *productRepository) UpdatePromoProduct(productID string, dto dto.PromoProductDTO) error {
	product, err := r.GetProductByID(productID)
	if err != nil {
		return err
	}

	product.DiscountStartDate = dto.DiscountStartDate
	product.DiscountEndDate = dto.DiscountEndDate
	product.DiscountPercentage = dto.DiscountPercentage
	product.IsPromo = dto.IsPromo

	if err := r.db.Save(&product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) Save(product *Product) (*Product, error) {
	if err := r.db.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) GetRecommendedProducts() ([]Product, error) {
	var products []Product
	err := r.db.Where("is_recommended = ?", true).Find(&products).Error
	return products, err
}

func (r *productRepository) FindAll(name string, limit, offset int) ([]Product, int64, error) {
	var products []Product
	var total int64

	query := r.db.Model(&Product{})
	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	err := query.Count(&total).Limit(limit).Offset(offset).Find(&products).Error
	return products, total, err
}

func (r *productRepository) GetProductsByCategory(categoryID string, limit, offset int) ([]Product, error) {
	log.Printf("Executing query with categoryID: %s, limit: %d, offset: %d", categoryID, limit, offset)
	var products []Product
	query := `
			SELECT p.id, p.name, p.description, p.price, p.discount, p.image, p.is_new, p.rating, p.total_rating, 
						 p.category_id, c.name AS category_name
			FROM products p
			JOIN categories c ON p.category_id = c.id
			WHERE c.id = $1
			LIMIT $2 OFFSET $3`

	if err := r.db.Raw(query, categoryID, limit, offset).Scan(&products).Error; err != nil {
		return nil, err
	}

	for i := range products {
		if products[i].Category.Name == "" {
			products[i].Category.Name = "Unknown"
		}
		products[i].Category.ID = categoryID
	}

	return products, nil
}
