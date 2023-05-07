package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("Index is %v and Value is %v\n", index, day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto label
		}

		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("Value is :", rogueValue)
		rogueValue++
	}

label:
	fmt.Println("You are using goto label")
}
