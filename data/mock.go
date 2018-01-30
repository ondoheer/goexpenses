package data

import (
	"../models"
)

var Categories []models.Category

func PopulateCategories() {
	Categories = append(Categories, models.Category{ID: 3, Label: "Compras", Name: "compras", User: 1})
	Categories = append(Categories, models.Category{ID: 2, Label: "Fiesta", Name: "fiesta", User: 1})
	Categories = append(Categories, models.Category{ID: 1, Label: "Transporte", Name: "transporte", User: 1})
	Categories = append(Categories, models.Category{ID: 4, Label: "Viajes", Name: "viajes", User: 1})

}
