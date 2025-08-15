package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods tutorial.")

	akshat := User{"Akshat", "0xDarkXnight@joker.bat", true, 18}
	fmt.Println(akshat)
	fmt.Printf("Akshat details are: %+v\n", akshat)
	fmt.Printf("Name is %v and Email is %v.\n", akshat.Name, akshat.Email)
	akshat.GetStatus()
	akshat.NewEmail()
	fmt.Printf("Name is %v and Email is %v.\n", akshat.Name, akshat.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewEmail() {
	u.Email = "dark@joker.bat"
	fmt.Println("Email of this user is:", u.Email)
}
