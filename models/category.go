package models

type Category struct {
	ID     int    `json:"id,omitempty"`
	Label  string `json:"label,omitempty"`
	Name   string `json:"name,omitempty"`
	UserID int    `json:"userId,omitempty"`
}
