package main

import "fmt"

type Vehicle interface {
	Start()
	Stop()
}

type Car struct {
	Brand string
}

func (c Car) Start() {
	fmt.Printf("%s car is starting. \n", c.Brand)
}

func (c Car) Stop() {
	fmt.Printf("%s car is stopping. \n", c.Brand)
}

func StartVehicle(v Vehicle) {
	v.Start()
}

func StopVehicle(v Vehicle) {
	v.Stop()
}

type Bike struct {
	Brand string
}

func (b Bike) Start() {
	fmt.Printf("%s bike is starting.\n", b.Brand)
}

func (b Bike) Stop() {
	fmt.Printf("%s bike is stopping. \n", b.Brand)
}

type Boat struct {
	Brand string
}

func (b Boat) Start() {
	fmt.Printf("%s boat is starting.\n", b.Brand)
}

func (b Boat) Stop() {
	fmt.Printf("%s boat is stopping.\n", b.Brand)
}

func main() {

	car := Car{Brand: "Toyota"}
	bike := Bike{Brand: "Harley"}

	StartVehicle(car)
	StartVehicle(bike)
}

// type Car struct {
// 	Brand string
// }

// func (c Car) Start() {
// 	fmt.Printf("%s car is starting.\n", c.Brand)
// }

// func (c Car) Stop() {
// 	fmt.Printf("%s car is stopping.\n", c.Brand)
// }

// type Bike struct {
// 	Brand string
// }

// func (b Bike) Start() {
// 	fmt.Printf("%s bike is starting.\n", b.Brand)
// }

// func (b Bike) Stop() {
// 	fmt.Printf("%s bike is stopping.\n", b.Brand)
// }

// func main() {
// 	car := Car{Brand: "Toyota"}
// 	bike := Bike{Brand: "Harley"}

// 	car.Start()
// 	car.Stop()

// 	bike.Start()
// 	bike.Stop()
// }
