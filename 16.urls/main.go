package main

import (
	"fmt"
	"net/url"
)

const URL = "http://dev.test:3030/learn?coursename=go&paymentid=12345"

func main() {
	fmt.Println("Welcome to url")
	fmt.Println(URL)
	requestUrl, err := url.Parse(URL)
	if err != nil {
		panic(err)
	}

	fmt.Println(requestUrl.Scheme)
	fmt.Println(requestUrl.Host)
	fmt.Println(requestUrl.Path)
	fmt.Println(requestUrl.Port())
	fmt.Println(requestUrl.RawQuery)
	qparams := requestUrl.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println("Single params", qparams["coursename"])

	for _, value := range qparams {
		fmt.Println("Params are: ", value)
	}

	partOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hasan",
	}

	anotherUrl := partOfUrl.String()
	fmt.Println("Another Url: ", anotherUrl)
}
