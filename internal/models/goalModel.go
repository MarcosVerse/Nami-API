package models

import "time"

type Goal struct {
	ID           uint      `json:"id"            gorm:"primaryKey"`
	UserID       uint      `json:"user_id"`
	Title        string    `json:"title"`
	TargetValue  float64   `json:"target_value"`
	CurrentValue float64   `json:"current_value"`
	CreatedAt    time.Time `json:"created_at"`
}
