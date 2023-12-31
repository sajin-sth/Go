package main

import "fmt"

const LoginToken string = "flewfjewpo" //public variable must be started with capital letter

func main() {
	var username string = "Sajin"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.4553927498472
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable id of type: %T \n", anotherVariable)

	// implicit type
	var website = "learncodeonine.in"
	fmt.Println(website)

	// no var style
	numberOfUSer := 300000
	fmt.Println(numberOfUSer)

	// usage of public variable
	fmt.Println(LoginToken)
	fmt.Printf("Variable id of type: %T \n", LoginToken)
}
