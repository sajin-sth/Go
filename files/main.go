package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This need to go in a file - LearnCodeONline.in"

	file, err := os.Create("./mylcogofile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./mylcogofile.txt")

}

func readFile(filename string) {
	data, err := os.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside the file is: ", string(data))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
