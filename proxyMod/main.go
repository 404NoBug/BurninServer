package main

import "fmt"

type Car struct {
}

type Vehicle interface {
	Drive()
}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	vehicle Vehicle
	driver  *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 16 {
		c.vehicle.Drive()
	} else {
		fmt.Println("Driver too young!")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{&Car{}, driver}
}

func main() {
	carProxy1 := NewCarProxy(&Driver{Age: 14})
	carProxy1.Drive()
	carProxy2 := NewCarProxy(&Driver{Age: 16})
	carProxy2.Drive()
}
