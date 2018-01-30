package models

import (
	"fmt"
)

type Expense struct {
	ID       int     `json:"id,omitempty"`
	Date     string  `json:"date,omitempty"`
	Name     string  `json:"name,omitempty"`
	Amount   float32 `json:"amount,omitempty"`
	User     int     `json:"user,omitempty"`
	Category int     `json:"category,omitempty"`
}

func (Expense) TableName() string {
	return "expense"
}

func (e Expense) String() string {
	return fmt.Sprintf("Expense< id: %d name: %s amount: %d date: %v for user: %v>", e.ID, e.Name, e.Amount, e.Date, e.User)
}
