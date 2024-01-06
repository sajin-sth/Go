package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)
	checkError(err)
	fmt.Printf("Response is of type %T\n", response)

	defer response.Body.Close() //caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)
	checkError(err)
	content := string(databytes)
	fmt.Println(content)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}