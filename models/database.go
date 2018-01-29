package models

import "github.com/go-pg/pg"

func GetDB() *pg.DB {

	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "",
		Database: "postgres",
	})

	return db
}
