package main

import (
	"fmt"
)

// Public
const LoginToken string = "token"

func main() {
	var username string = "Ela"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type : %T \n", smallValue)

	var smallFloat float32 = 255.56787897
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	var bigFloat float64 = 255.56787897
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type : %T \n", bigFloat)

	// Default values and some aliases
	var anotherValue int
	fmt.Println(anotherValue)
	fmt.Printf("Variable is of type : %T \n", anotherValue)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type : %T \n", anotherString)

	// Implicit type
	var website = "github"
	fmt.Println(website)

	// No var style
	num := 30000
	fmt.Println(num)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)
}
