package main

import "fmt"

type vehicle interface {
	speed() int
}

type car struct {
	maxSpeed int
}
type bike struct {
	maxSpeed int
}

func (c car) speed() int {
	return c.maxSpeed
}
func (b bike) speed() int {
	return b.maxSpeed
}

func main() {
	var v vehicle
	var mycar car
	fmt.Print("car max speed:")
	fmt.Scan(&mycar.maxSpeed)

	var mybike bike
	fmt.Print("bike max speed:")
	fmt.Scan(&mybike.maxSpeed)

	v = mycar
	fmt.Println("car speed:", v.speed())

	v = mybike
	fmt.Println("bike speed:", v.speed())

}
