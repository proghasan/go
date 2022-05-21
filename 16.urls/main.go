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
}
