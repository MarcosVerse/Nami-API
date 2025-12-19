package goal

type UpdateGoalDTO struct {
	Title        *string  `json:"title,omitempty"`
	TargetValue  *float64 `json:"target_value,omitempty"`
	CurrentValue *float64 `json:"current_value,omitempty"`
}
