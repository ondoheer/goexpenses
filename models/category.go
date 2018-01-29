package models

import (
	"errors"
	"fmt"
)

type Category struct {
	ID     int    `json:"id,omitempty"`
	Label  string `json:"label,omitempty"`
	Name   string `json:"name,omitempty"`
	UserID int    `json:"userId,omitempty"`
}

func (Category) TableName() string {
	return "category"
}

func (c Category) String() string {
	return fmt.Sprintf("Category<%d %s>", c.ID, c.Label)
}

func (c Category) GetById(id int) (Category, error) {
	db := GetDB()

	category := Category{ID: id}

	if db.First(&category, id).RecordNotFound() {

		return category, errors.New("Item no encontrado")
	}

	return category, nil
}
