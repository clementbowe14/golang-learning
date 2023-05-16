package models

import "fmt"

type Item interface {
	GetPrice() float64
	GetWeight() float64
}

type Customer struct {
	name  string
	items []Item
}

type Bike interface {
	Item
	GetWheelCount() int32
	Ride()}

type MotorBike struct {
	wheelCount int32
	maxSpeed   float64
	model      string
	price      float64
	weight     float64
}

type Scooter struct {
	wheelCount int32
	model      string
	price      float64
	maxSpeed   float64
	isStanding bool
	isElectric bool
	weight     float64
}

type Bicycle struct {
	wheelCount        int32
	model             string
	price             float64
	hasGears          bool
	hasTrainingWheels bool
	weight            float64
}



func (m MotorBike) GetWheelCount() int32 {
	return m.wheelCount
}

func (m MotorBike) Ride() {
	fmt.Println("Vroom! I am riding on a motorcycle")
}

func (m MotorBike) GetMaxSpeed() float64 {
	return m.maxSpeed
}

func (m MotorBike) GetPrice() float64 {
	return m.price
}

func (m MotorBike) GetWeight() float64 {
	return m.weight
}

func (m MotorBike) GetModelName() string {
	return m.model
}
func (s Scooter) GetWheelCount() int32 {
	return s.wheelCount
}

func (s Scooter) GetMaxSpeed() float64 {
	if s.isElectric {
		return s.maxSpeed
	}

	return 0
}

func (s Scooter) IsStandingScooter() bool {
	return s.isStanding
}

func (s Scooter) IsElectric() bool {
	return s.isElectric
}

func (s Scooter) GetModelName() string {
	return s.model
}

func (s Scooter) Ride() {
	fmt.Println("Beep Beep! Scooter coming through")
}

func (s Scooter) GetWeight() float64 {
	return s.weight
}

func (b Bicycle) GetWheelCount() int64 {
	return b.GetWheelCount()
}

func (b Bicycle) GetModelName() string {
	return b.model
}

func (b Bicycle) GetPrice() float64 {
	return b.price
}

func (b Bicycle) HasGears() bool {
	return b.hasGears
}

func (b Bicycle) HasTrainingWheels() bool {
	return b.hasTrainingWheels
}

func (b Bicycle) GetWeight() float64 {
	return b.weight
}