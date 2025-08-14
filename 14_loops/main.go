package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for i, day := range days {
		fmt.Printf("index is %v and value is %v\n", i, day)
	}

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 2 {
			goto bat
		}

		if rogueValue == 5 {
			rogueValue++
			continue
		}

		fmt.Println("Value is:", rogueValue)
		rogueValue++
	}

bat:
	fmt.Println("El Rata Elada!")
}
