package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMess := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro message is: ", myMess)
}

func greeter() {
	fmt.Println("Namaste from golang")
}

func adder(x int, y int) int {
	return (x + y)
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi pro result"
}
