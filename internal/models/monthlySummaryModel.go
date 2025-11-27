package models

import "time"

type MonthlySummary struct {
	ID                  uint      `json:"id"                   gorm:"primaryKey"`
	UserID              uint      `json:"user_id"`
	Year                int       `json:"year"`
	Month               int       `json:"month"`
	TotalIncome         float64   `json:"total_income"`
	TotalExpense        float64   `json:"total_expense"`
	CategoriesBreakdown string    `json:"categories_breakdown"`
	CreatedAt           time.Time `json:"created_at"`
}
