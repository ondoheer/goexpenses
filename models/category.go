package models

import (
	"fmt"
)

type Category struct {
	ID    int    `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Name  string `json:"name,omitempty"`
	User  int    `json:"user,omitempty"`
}

func (Category) TableName() string {
	return "category"
}

func (c Category) String() string {
	return fmt.Sprintf("Category< id: %d label: %s >", c.ID, c.Label)
}
