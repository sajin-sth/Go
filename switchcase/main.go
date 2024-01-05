package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	randomMinutes := rng.Intn(15) + 1
	fmt.Printf("Random time: %v minutes\n", randomMinutes)

	// rand.Seed(time.Now().UnixNano()) ->this method is deprecated
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can open or move 1 spot.")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot and roll dice again")
		fallthrough
	default:
		fmt.Println("Congratulation")
	}
}
