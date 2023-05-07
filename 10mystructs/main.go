package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// No inheritance in golang; No super or parent
	user := User{"Jane", "jane90@gmail.com", true, 25}
	fmt.Println(user)
	fmt.Printf("User details are : %+v\n", user)
	fmt.Printf("Name is %v and Email is %v\n", user.Name, user.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
