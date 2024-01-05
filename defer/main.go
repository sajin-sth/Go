package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two") //-> last in first out (world, one, two), so two gets printed out first
	fmt.Println("Hello")

	myDefer() //->this gets printed first because it hasnot yet recognize other defer statements
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
