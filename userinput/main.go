package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) // reading/taking user input
	fmt.Print("enter the rating for our pizza: ")

	// comma ok | error syntax just like try_catch syntax
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Thanks for rating,", input)
		fmt.Printf("Type of this rating is %T\n", input)
	}

	/*
		## Another way to get user input

		var userInput string
		_ , err := fmt.Scan(&userInput)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Thanks for rating,", userInput)
		fmt.Printf("Type of this rating is %T\n", userInput)
		}
	*/
}
