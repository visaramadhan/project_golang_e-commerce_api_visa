package category

import "time"

type Category struct {
	ID        string    `gorm:"type:uuid;primaryKey;not null;unique" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
