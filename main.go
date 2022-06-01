package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	// fmt.Println("There is a sign near the entrance that reads 'No Minors'.")
	// var age = 41
	// var minor = age < 18
	// fmt.Printf("At age %v, am I a minor? %v\n", age, minor)

	// var command = "go east"

	// if command == "go east" {
	// 	fmt.Println("You head further up the mountain.")
	// } else if command == "go inside" {
	// 	fmt.Println("You enter the cave where you live out the rest of your life.")
	// } else {
	// 	fmt.Println("Didn't quite get that.")
	// }

	// var room = "lake"
	// switch {
	// case room == "cave":
	// 	fmt.Println("You find yourself in a dimly lit cavern.")
	// case room == "lake":
	// 	fmt.Println("The ice seems solid enough.")
	// 	fallthrough
	// case room == "underwater":
	// 	fmt.Println("The water is freezing cold.")
	// }

	// for count := 10; count > 0; {
	// 	fmt.Println(count)
	// 	time.Sleep(time.Second)
	// 	count--
	// }
	// fmt.Println("Liftoff!")

	// switch num := rand.Intn(10); num {
	// case 0:
	// 	fmt.Println("Space Adventures")
	// case 1:
	// 	fmt.Println("SpaceX")
	// case 2:
	// 	fmt.Println("Virgin Galactic")
	// default:
	// 	fmt.Println("Random spaceline #", num)
	// }

	for count := 0; count <= 10; count++ {

		year := 2000 + rand.Intn(22)
		month := rand.Intn(12) + 1
		daysInMonth := 31

		switch month {
		case 2:
			if year%4 == 0 {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}
		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}
}
