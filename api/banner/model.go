package banner

import "time"

type Banner struct {
	ID        string    `gorm:"type:uuid;primaryKey;not null;unique" json:"id"`
	Photo     string    `gorm:"type:varchar(255)" json:"photo"`
	Title     string    `gorm:"type:varchar(255)" json:"title"`
	Subtitle  string    `gorm:"type:varchar(255)" json:"subtitle"`
	PathPage  string    `gorm:"type:varchar(255)" json:"path_page"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
