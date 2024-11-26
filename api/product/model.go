package product

import (
	"time"

	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/category"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string  `gorm:"type:uuid;primaryKey;not null;unique" json:"id"`
	Name        string  `gorm:"type:varchar(255);not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Discount    float64 `gorm:"type:decimal(5,2)" json:"discount"`
	// FinalPrice  float64        `gorm:"type:decimal(10,2);not null" json:"finalPrice"`
	Stock       int            `gorm:"type:int;not null" json:"stock"`
	Image       string         `gorm:"type:varchar(255)" json:"image"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Rating      float64        `gorm:"type:decimal(3,2)" json:"rating"`
	TotalRating int            `json:"total_rating"`

	CategoryID string            `json:"category_id"`
	Category   category.Category `json:"category"`

	DiscountStartDate  time.Time `json:"discount_start_date"`
	DiscountEndDate    time.Time `json:"discount_end_date"`
	DiscountPercentage float64   `json:"discount_percentage"`
	IsPromo            bool      `json:"is_promo"`

	IsRecommended bool `json:"is_recommended"`
}

type ProductDetail struct {
	ID          string      `json:"id" gorm:"primaryKey"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Rating      float64     `json:"rating"`
	TotalRating int         `json:"total_rating"`
	Variations  []Variation `json:"variations" gorm:"foreignKey:ProductID"`
	Images      []string    `json:"images" gorm:"foreignKey:ProductID"`
}

type Variation struct {
	ID        string `json:"id" gorm:"primaryKey"`
	ProductID string `json:"product_id" gorm:"index"`
	Color     string `json:"color"`
	Size      string `json:"size"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}
