package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proResult, myMessage := proAdder(2, 5, 7, 4, 8)
	fmt.Println("Pro Result is:", proResult)
	fmt.Println("Pro Message is:", myMessage)
}

func adder(a, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hi, from Pro result function"
}

func greeter() {
	fmt.Println("Namastey from Golang")
}
