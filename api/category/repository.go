package category

import (
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]Category, error)
	GetCategoryByID(id string) (Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) CategoryRepository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Category, error) {
	var categories []Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *repository) GetCategoryByID(id string) (Category, error) {
	var category Category
	err := r.db.First(&category, "id = ?", id).Error
	if err != nil {
		return category, err
	}
	return category, nil
}
