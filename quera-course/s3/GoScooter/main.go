package main

import "fmt"

type Car struct {
	// TODO
    speed int
	battery int
}

func NewCar(speed, battery int) *Car {
	// TODO
	var car Car
	car.speed = speed
	car.battery = battery
	return &car
}
func GetSpeed(car *Car) int {
	// TODO
	return car.speed
}
func GetBattery(car *Car) int {
	// TODO
	return car.battery
}
func ChargeCar(car *Car, minutes int) {
	// TODO
	inc := minutes/2
	car.battery += inc
	if car.battery > 100 {
	 car.battery = 100   
	}
}
func TryFinish(car *Car, distance int) string {
	// TODO
	if car.battery*2 < distance{
	    car.battery = 0
	    return ""
	}
	car.battery -= distance/2
	timeElapsed := float32(distance)/float32(car.speed)
	r := fmt.Sprintf("%.2f",timeElapsed)
	return r
}

