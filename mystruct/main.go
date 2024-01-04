package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")

	sajin := User{"Sajin", "shresthasajin59@gmail.com", true, 20}
	fmt.Println(sajin)
	fmt.Printf("Sajin details are: %+v\n", sajin)
	fmt.Printf("Name is %v and email is %v\n", sajin.Name, sajin.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
