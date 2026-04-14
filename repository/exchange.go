package repository

import (
	"time"

	"gorm.io/gorm"
)

type ExchangeRate struct{
	gorm.Model
	BaseCurrency string `json:"base_currency" binding:"required"`
	TargetCurrency string `json:"target_currency" binding:"required"`	
	Rate float64 `json:"rate" binding:"required"`
	DateTime time.Time `json:"date_time"`
}