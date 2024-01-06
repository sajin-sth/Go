package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=iewjd9" 

func main() {
	fmt.Println("Welcome to handeling url in golang")
	fmt.Println(myurl)

	//parsing
	result, err := url.Parse(myurl)
	checkError(err)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])
	for _, val := range(qparams) {
		fmt.Println("param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=hitesh",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}