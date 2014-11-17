package main

import "fmt"

// Simple example of aggregation
type Food struct {
	name string // Aggregated field
}

func (this *Food) eat() {
	fmt.Println("Eat: ", this.name)
}

// Demonstration of embedding and aggregation
type Recipe struct {
	Food               // Embedded field
	Ingredients []Food // Aggregated field
}

func (this *Recipe) eat() {
	fmt.Println("Eat: ", this.name, "\n----------------------")

	for _, ingredient := range this.Ingredients {
		ingredient.eat()
	}
}

func main() {
	fries := Food{
		name: "fries",
	}

	fries.eat()

	fmt.Println("\n------------------------------------------------------------------\n")

	cheeseburger := Recipe{
		Ingredients: []Food{{name: "bun"}, {name: "cucumber"}, {name: "cheese"}, {name: "cutlet"}, {name: "bun"}},
	}

	cheeseburger.Food.name = "Small cheeseburger"
	cheeseburger.Food.eat()

	fmt.Println("\n------------------------------------------------------------------\n")

	cheeseburger.name = "Big cheeseburger"
	cheeseburger.eat()
}

