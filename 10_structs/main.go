package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs tutorial.")

	akshat := User{"Akshat", "0xDarkXnight@gmail.com", true, 18}
	fmt.Println(akshat)
	fmt.Printf("Akshat details are: %+v\n", akshat)
	fmt.Printf("Name is %v and Email is %v.\n", akshat.Name, akshat.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
