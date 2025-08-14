package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "New User"
	} else if loginCount > 10 {
		result = "Regular User"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	number := 9

	if number%2 == 0 {
		fmt.Println("Number is odd")
	} else {
		fmt.Println("Number is even")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}
}
