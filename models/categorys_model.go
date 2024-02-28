package models

// Airline document model
type Categorys struct {
	title       string `json:"title,omitempty" example:"SAF"`
	description string `json:"description,omitempty" binding:"required" example:"United States"`
}
