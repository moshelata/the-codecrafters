package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("=== Simple Calculator ===")

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	// Check for invalid operator
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Error: Invalid operator! Use +, -, *, /")
		return
	}

	// Check for divide by zero
	if operator == "/" && num2 == 0 {
		fmt.Println("Error: Cannot divide by zero!")
		return
	}

	// Do the math
	var result float64
	if operator == "+" {
		result = num1 + num2
	} else if operator == "-" {
		result = num1 - num2
	} else if operator == "*" {
		result = num1 * num2
	} else if operator == "/" {
		result = num1 / num2
	}

	fmt.Println("Result:", num1, operator, num2, "=", result)
}
