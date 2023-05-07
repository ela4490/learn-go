package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is :", result)

	proResult, myMessage := proAdder(2, 5, 8, 7, 3)
	fmt.Println("Pro result is :", proResult)
	fmt.Println("Pro message is :", myMessage)
}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}

	return total, "Uses proAdder function"
}

func greeter() {
	fmt.Println("Namstey from golang")
}
