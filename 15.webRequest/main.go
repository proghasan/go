package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://jsonplaceholder.typicode.com/todos"

func main() {
	fmt.Println("Welcome web request")
	response, err := http.Get(URL)
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
