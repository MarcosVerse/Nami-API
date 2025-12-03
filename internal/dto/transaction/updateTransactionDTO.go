package transaction

type UpdateTransactionDTO struct {
	Category *string  `json:"category,omitempty"`
	Value    *float64 `json:"value,omitempty"`
	Type     *string  `json:"type,omitempty"`
}
