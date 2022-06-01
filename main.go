package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")
	var age = 41
	var minor = age < 18
	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)

	var command = "go east"

	if command == "go east" {
		fmt.Println("You head further up the mountain.")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else {
		fmt.Println("Didn't quite get that.")
	}

	var room = "lake"
	switch {
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case room == "underwater":
		fmt.Println("The water is freezing cold.")
	}

	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Liftoff!")
}
