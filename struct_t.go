package main

import "fmt"

type Car struct {
	Make           string
	Model          string
	Year           int
	Mileage        int
	PreviousOwners int
}

func (c Car) DisplayInfo() {
	fmt.Println("The car is", c.Model, "made by", c.Make, "in the year", c.Year, "with mileage", c.Mileage, "has", c.PreviousOwners, "previous owners.")
}
func (c *Car) Drive(distance int) {
	c.Mileage += distance
}

func (c *Car) ChangeYear(newYear int) {
	c.Year = newYear
}

func main() {
	myCar := Car{
		Make:           "Honda",
		Model:          "Civic",
		Year:           2020,
		Mileage:        15000,
		PreviousOwners: 1,
	}
	myCar.DisplayInfo()
	myCar.Drive(300)
	myCar.ChangeYear(2023)
	myCar.DisplayInfo()
}
