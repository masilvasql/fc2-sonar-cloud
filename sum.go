package main

import "fmt"

func main() {
	fmt.Println(sum(2, 2))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mult(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func mod(a int, b int) int {
	return a % b
}

func pow(a int, b int) int {
	return a ^ b
}
