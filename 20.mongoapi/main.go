package main

import (
	"fmt"
	"github.com/proghasan/mongoapi/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Mongodb API")
	fmt.Println("Server is getting starting...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening server 4000")
}
