package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=123445445"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	// Parsing
	result, _ := url.Parse(myurl)
	fmt.Println("Scheme :", result.Scheme)
	fmt.Println("Host :", result.Host)
	fmt.Println("Path :", result.Path)
	fmt.Println("Raw Query :", result.RawQuery)
	fmt.Println("Port :", result.Port())

	qparams := result.Query()
	fmt.Printf("The type of query params are : %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is :", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=jane",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Another url :", anotherUrl)
}
