package main

import "fmt"

//Car structs and map
type car struct {
	price int
	brand carBrand
}

type carBrand struct{
	manufacturer string
	model string
}

func (i *carBrand) getCarName() string {
	return i.manufacturer + " " + i.model
}


func printCarMap(){
	for i,j := range carMap{
		fmt.Println(i, j.brand.getCarName(), "\t\tPrice: $",j.price)
	}
}

func sellCar(){
	var n int
	fmt.Print("Which car would you like to sell? ")
	fmt.Scan(&n)
	fmt.Println("You have sold the", carMap[n].brand.manufacturer, carMap[n].brand.model, "for $", carMap[n].price)
	deleteCar(n)
}

func buyCar() {
	var p int
	fmt.Print("Enter the price of the car: ")
	fmt.Scan(&p)
	var manu string
	fmt.Print("Enter the manufacturer of the car: ")
	fmt.Scan(&manu)
	var mod string
	fmt.Print("Enter the model of the car: ")
	fmt.Scan(&mod)
	addCar(p, manu, mod)
}

func deleteCar(n int) {
	delete(carMap, n)
	fmt.Println("Car", n, "has been successfully deleted")
	printCarMap()
}

func addCar(priceCar int, carManufacturer, carModel string){
	newInt := len(carMap) + 1
	carMap[newInt] = car{price: priceCar, brand: carBrand{manufacturer: carManufacturer, model: carModel}}
	fmt.Println("Car", newInt, "has been succesfully added")
	printCarMap()
}

var carMap map[int]car = map[int]car{1: car1, 2: car2, 3: car3, 4: car4, 5: car5}
var car1 car = car{price: 25000, brand: carBrand{manufacturer: "Toyota", model: "Camry"}}
var car2 car = car{price: 30000, brand: carBrand{manufacturer: "Honda", model: "Pilot"}}
var car3 car = car{price: 46000, brand: carBrand{manufacturer: "Mercedes", model: "CLA"}}
var car4 car = car{price: 28000, brand: carBrand{manufacturer: "Dodge", model: "Challenger"}}
var car5 car = car{price: 35000, brand: carBrand{manufacturer: "Subaru", model: "WRX"}}


//Substring Section
var phrase string = "Hello, World"

func sliceSub(start , end int) []rune{
	var tempSubstring []rune = []rune(phrase)
	var tempSlice []rune = tempSubstring[2:9]
	var newSubstring []rune = tempSlice[start:end]
	return newSubstring
}

//Main
func main(){
	//Part 1
	printCarMap()
	fmt.Println()
	// sellCar()
	deleteCar(3)
	// printCarMap()
	// buyCar()
	addCar(20000, "Ford", "Fusion")
	printCarMap()
	//Use of the substring function (part 2)
	var substring []rune = sliceSub(1,4)
	fmt.Println(string(substring))
}