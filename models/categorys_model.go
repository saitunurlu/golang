package models

// Airline document model
type Categorys struct {
	Title       string `json:"title,omitempty" example:"SAF"`
	Description string `json:"description,omitempty" binding:"required" example:"United States"`
}
