package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const URL = "https://jsonplaceholder.typicode.com"

func main() {
	fmt.Println("Welcome to request Get/Post")

	// get all record
	//PerformGetRequest()

	// insert new record
	//PerformPostJsonRequest()

	// form data
	PerformPostFormRequest()
}

func PerformGetRequest() {

	// get request call
	response, err := http.Get(URL + "/todos/")
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println("Response status is: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is:", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	requestBody := strings.NewReader(`
		{
			title: "Hasan",
			body: "Your body goes to here",
			userId: 1,
		}
	`)

	response, err := http.Post(URL+"/posts", "application/json", requestBody)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	// formdata
	data := url.Values{}
	data.Add("title", "New title")
	data.Add("body", "Body goes here")
	data.Add("userId", "1")

	response, err := http.PostForm(URL+"/posts", data)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
