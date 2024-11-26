package address

type Address struct {
	ID       string `json:"id" gorm:"primaryKey"`
	UserID   string `json:"user_id"`
	Street   string `json:"street"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
	Default  bool   `json:"default"`
}
