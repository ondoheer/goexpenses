package models

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int        `json:"id,omitempty"`
	Username string     `json:"username,omitempty"`
	Name     string     `json:"name,omitempty"`
	Email    string     `json:"email,omitempty"`
	Password []byte     `json:"password,omitempty"`
	Expenses []Expense  `json:"expenses,omitempty"`
	Category []Category `json:"categories,omitempty"`
}

func (User) TableName() string {
	return "user"
}

func (u User) String() string {
	return fmt.Sprintf("User< id: %d name: %s username: %s email: %s >", u.ID, u.Name, u.Username, u.Email)
}

func (u User) HashPassword(plainPassword []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}
	return hash

}
