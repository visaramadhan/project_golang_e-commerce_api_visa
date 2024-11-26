package users

type Users struct {
	ID          string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	Name        string `gorm:"type:varchar(255);not null" json:"name" binding:"required,alphanum"`
	Email       string `gorm:"type:varchar(255);not null" json:"email,omitempty" binding:"required,email"`
	PhoneNumber string `gorm:"type:varchar(255);not null" json:"phoneNumber,omitempty" binding:"required,numeric"`
	Password    string `gorm:"type:varchar(255);not null" json:"password,omitempty" binding:"required,alphanum,min=6"`
}
