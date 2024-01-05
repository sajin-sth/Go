package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")

	person := User{"Sajin", "shresthaperson59@gmail.com", true, 20}
	fmt.Println(person)
	fmt.Printf("person details are: %+v\n", person)
	fmt.Printf("Name is %v and email is %v\n", person.Name, person.Email)

	person.GetStatus()

	person.NewMail()
	fmt.Println(person.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
