package main

//package GasMileage

import "fmt"

func main() {
	var num1 int = 0
	var num2 = 0

	fmt.Print("Enter number of miles driven: ")
	fmt.Scan(&num1)

	fmt.Print("Enter number of gallon: ")
	fmt.Scan(&num2)

	if num2 != 0 {
		result := float64(num1) / float64(num2)
		fmt.Printf("The result of number of miles per gallon %f miles", result)
	} else {
		fmt.Println("Division by zero is disallowed.")
	}
}
