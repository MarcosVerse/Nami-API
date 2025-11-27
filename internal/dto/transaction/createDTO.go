package transaction

type CreateDTO struct {
	Category string  `json:"category"`
	Value    float64 `json:"value"`
	Type     string  `json:"type"`
	UserID   uint    `json:"user_id"`
}
