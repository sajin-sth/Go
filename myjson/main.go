package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct to represent your data
type Person struct {
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	City       string   `json:"city"`
	Email      []string `json:"email,omitempty"`
	SecretCode string   `json:"-"`
}

func main() {
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	// Create an instance of the struct
	person := []Person{
		{"Sajin Shrestha", 20, "Kathmandu", []string{"shresthasajin59@gmail.com", "sajinxtha1@gmail.com"}, "abc123"},
		{"John Doe", 30, "New York", nil, "def123"},
	}
	// person := Person{
	// 	Name:  "John Doe",
	// 	Age:   30,
	// 	City:  "New York",
	// 	Email: "john.doe@example.com",
	// }

	// Convert the struct to JSON
	jsonData, err := json.MarshalIndent(person, "", "\t")
	if err != nil {
		panic(err)
	}

	// Print the JSON data
	fmt.Println(string(jsonData))
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"name": "Sajin Shrestha",
		"age": 20,
		"city": "Kathmandu",
		"email": ["shresthasajin59@gmail.com","sajinxtha1@gmail.com"]
	}
	`)

	var personData Person
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &personData)
		fmt.Printf("%#v\n", personData)
	} else {
		fmt.Println("JSON as not valid")
	}

	fmt.Println("")

	//some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	// fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("KEY:%v and VALUE:%v and TYPE:%T\n", k, v, v)
	}

}
