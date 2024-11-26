package banner

import (
	"gorm.io/gorm"
)

type BannerRepository interface {
	GetAllBanners() ([]Banner, error)
}

type bannerRepository struct {
	db *gorm.DB
}

func NewBannerRepository(db *gorm.DB) BannerRepository {
	return &bannerRepository{db: db}
}

func (r *bannerRepository) GetAllBanners() ([]Banner, error) {
	var banners []Banner
	err := r.db.Find(&banners).Error
	return banners, err
}
