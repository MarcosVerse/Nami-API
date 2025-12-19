package goal

type CreateGoalDTO struct {
	Title       string  `json:"title"        binding:"required"`
	TargetValue float64 `json:"target_value" binding:"required"`
	UserID      uint    `json:"user_id"      binding:"required"`
}
