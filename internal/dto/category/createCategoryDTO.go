package category

type CreateCategoryDTO struct {
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	UserID uint   `json:"user_id"`
}
