package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const lightSpeed = 299792 // km/s
	var distance = 56000000 // km
	fmt.Println(distance/lightSpeed, "seconds")
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	var random = rand.Int()

	fmt.Println("Random number: ", random)
}