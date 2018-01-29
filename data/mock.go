package data

import (
	"../models"
)

var Categories []models.Category

func PopulateCategories() {
	Categories = append(Categories, models.Category{ID: 3, Label: "Compras", Name: "compras", UserID: 1})
	Categories = append(Categories, models.Category{ID: 2, Label: "Fiesta", Name: "fiesta", UserID: 1})
	Categories = append(Categories, models.Category{ID: 1, Label: "Transporte", Name: "transporte", UserID: 1})
	Categories = append(Categories, models.Category{ID: 4, Label: "Viajes", Name: "viajes", UserID: 1})

}
