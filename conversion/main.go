package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza shop")
	fmt.Print("Please rate our pizza between 1 to 5: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 32)
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("thanks for rating,", input)
		fmt.Println("Added 1 to your rating:", numRating + 1)
	}
}
