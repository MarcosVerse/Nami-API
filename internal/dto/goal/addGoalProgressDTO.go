package goal

type AddProgressDTO struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}
