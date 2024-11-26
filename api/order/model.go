package order

import (
	"time"
)

type Order struct {
	ID        string    `gorm:"type:uuid;primaryKey;not null;unique" json:"id"`
	ProductID string    `gorm:"type:uuid;not null;" json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
