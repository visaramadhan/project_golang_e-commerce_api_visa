package seeder

import (
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/banner"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

func SeedBanners(db *gorm.DB) error {
	banners := []banner.Banner{
		{
			ID:       uuid.New().String(),
			Photo:    "banner1.jpg",
			Title:    "Sale",
			Subtitle: "50% off on all items",
			PathPage: "/sale",
		},
		{
			ID:       uuid.New().String(),
			Photo:    "banner2.jpg",
			Title:    "New Arrivals",
			Subtitle: "Check out our latest products",
			PathPage: "/new-arrivals",
		},
		{
			ID:       uuid.New().String(),
			Photo:    "banner3.jpg",
			Title:    "Exclusive Deals",
			Subtitle: "Special prices for members",
			PathPage: "/deals",
		},
	}

	for _, banner := range banners {
		if err := db.Create(&banner).Error; err != nil {
			return err
		}
	}
	return nil
}
