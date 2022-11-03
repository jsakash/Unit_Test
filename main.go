package main

import "fmt"

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func Multiply(x int) (result int) {
	result = x * 10
	return result
}

func main() {
	result := Calculate(2)
	fmt.Println("The Result is := ", result)

}

func Square(x int) int {

	result := x * x
	return result
}
