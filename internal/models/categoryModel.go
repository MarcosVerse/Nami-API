package models

type Category struct {
	ID     uint   `json:"id"      gorm:"primaryKey"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Icon   string `json:"icon"` // talvez implementar
}
