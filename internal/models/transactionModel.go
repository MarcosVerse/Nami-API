package models

import "time"

type Transaction struct {
	ID        uint      `json:"id"         gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Type      string    `json:"type"`
	Category  string    `json:"category"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
