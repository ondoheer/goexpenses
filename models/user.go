package models

import (
	"fmt"
)

type User struct {
	ID       int        `json:"id,omitempty"`
	Username string     `json:"username,omitempty"`
	Name     string     `json:"name,omitempty"`
	Email    string     `json:"email,omitempty"`
	Password string     `json:"password,omitempty"`
	Expenses []Expense  `json:"expenses,omitempty"`
	Category []Category `json:"categories,omitempty"`
}

func (User) TableName() string {
	return "user"
}

func (u User) String() string {
	return fmt.Sprintf("User< id: %d name: %s username: %s email: %s >", u.ID, u.Name, u.Username, u.Email)
}
