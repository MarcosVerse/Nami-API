package category

type UpdateCategoryDTO struct {
	Name *string `json:"name,omitempty"`
	Icon *string `json:"icon,omitempty"`
}
