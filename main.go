package main

import "fmt"

func main() {
	fmt.Println("Calculator")
	fmt.Println("Choose first action: (+, -, /, *)")

	var action string
	fmt.Scan(&action)

	var a float64
	var b float64

	fmt.Println("Type first number:")
	fmt.Scan(&a)

	fmt.Println("Type second number:")
	fmt.Scan(&b)

	switch {
	case action == "+":
		fmt.Println("a + b = " + fmt.Sprint(a+b))
	case action == "-":
		fmt.Println("a - b = " + fmt.Sprint(a-b))
	case action == "/":
		fmt.Println("a / b = " + fmt.Sprint(a/b))
	case action == "*":
		fmt.Println("a * b = " + fmt.Sprint(a*b))
	}

}
