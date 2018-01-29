package models

import "fmt"

type Category struct {
	ID     int    `json:"id,omitempty"`
	Label  string `json:"label,omitempty"`
	Name   string `json:"name,omitempty"`
	UserID int    `json:"userId,omitempty"`
}

func (c Category) String() string {
	return fmt.Sprintf("Category<%d %s %v>", c.ID, c.Label)
}
