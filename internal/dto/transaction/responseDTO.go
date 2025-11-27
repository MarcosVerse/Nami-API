package transaction

type ResponseDTO struct {
	ID          uint    `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	UserID      uint    `json:"user_id"`
}
