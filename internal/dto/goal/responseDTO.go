package goal

type ResponseDTO struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	TargetValue  float64 `json:"target_value"`
	CurrentValue float64 `json:"current_value"`
	UserID       uint    `json:"user_id"`
}
