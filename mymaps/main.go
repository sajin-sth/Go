package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in golang")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("js shorts for: ", languages["js"])

	delete(languages, "rb")
	fmt.Println("List of all languages: ", languages)

	//loop through maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}