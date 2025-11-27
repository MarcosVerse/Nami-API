package goal

type UpdateDTO struct {
	Name         *string  `json:"name,omitempty"`
	TargetValue  *float64 `json:"target_value,omitempty"`
	CurrentValue *float64 `json:"current_value,omitempty"`
}
