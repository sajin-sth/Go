package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sajin-sth/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")

	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", router.Router()))

	fmt.Println("Listening at localhost:4000")
}
