package store

import (
	"fmt"

	models "github.com/clementbowe14/golang-learning/exercises/bikestore/models"
)

type Store interface {
	PurchaseBikes(item []models.Item)
	CalculateMaintenanceCost(b models.Bike) float64
	CalculateShippingCost(b models.Bike) float64
	ItemInStock(itemName string) bool
}

type BikeStore struct {
	name string
	items map[string]bool
	maintenanceFee float64
	shippingFee float64
	isCardValid bool
	isCashValid bool
	isApplePayValid bool
	isGiftCardValid bool
	
}

func (b BikeStore) PurchaseBikes(items []models.Item) {
	var totalPrice float64
	for i := range items {
		totalPrice += items[i].GetPrice()
	}

	fmt.Printf("Total Price:  %f", totalPrice)
}

func (b BikeStore) CalculateMaintenanceCost(bike models.Bike) float64 {
	return (.10 * bike.GetPrice()) + (bike.GetWeight() * b.maintenanceFee) 
}

func(b BikeStore) CalculateShippingCost(item models.Item) float64 {
	return item.GetWeight()/100
}




