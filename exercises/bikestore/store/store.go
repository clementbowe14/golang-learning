package store

import (
	models "github.com/clementbowe14/golang-learning/exercises/bikestore/models"
)

type StoreHandler interface {
	PurchaseBikes(item []models.Item)
	CalculateMaintenanceCost(b models.Bike) float64
	CalculateShippingCost(b models.Bike) float64
	ItemInStock(itemName string) bool
}

type BikeStoreHandler struct {

}



func (b BikeStoreHandler) PurchaseBikes(items []models.Item) {
	var totalPrice float64
	for i := range(items) {
		totalPrice += items[i].GetPrice()
	}

	fmt.Printf("Purchased %f tal")
}
