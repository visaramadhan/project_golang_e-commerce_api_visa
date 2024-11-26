package dto

import "time"

type PromoProductDTO struct {
	DiscountStartDate  time.Time `json:"discount_start_date"`
	DiscountEndDate    time.Time `json:"discount_end_date"`
	DiscountPercentage float64   `json:"discount_percentage"`
	IsPromo            bool      `json:"is_promo"`
}
