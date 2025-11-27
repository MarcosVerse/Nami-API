package summary

type ResponseDTO struct {
	ID           uint    `json:"id"`
	Month        string  `json:"month"`
	TotalIncome  float64 `json:"total_income"`
	TotalExpense float64 `json:"total_expense"`
	Balance      float64 `json:"balance"`
	UserID       uint    `json:"user_id"`
}
