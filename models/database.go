package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDB() *gorm.DB {

	db, err := gorm.Open("postgres", "host=localhost user=ondoheer dbname=expenses sslmode=disable")

	if err != nil {
		panic(err)
	}

	return db
}
