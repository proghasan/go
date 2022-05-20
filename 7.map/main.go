package main

import "fmt"

func main() {
	fmt.Println("Welcome to map")
	languages := make(map[string]string)

	languages["GO"] = "Golang Programming"
	languages["JS"] = "JavaScript Scripting Language"
	languages["PHP"] = "PHP Scripting Language"
	languages["PY"] = "Python Scripting Language"

	fmt.Println("Language List: ", languages)

	delete(languages, "PHP")

	fmt.Println("After Deleted of Language List: ", languages)

	fmt.Println()
	for _, value := range languages {
		fmt.Println("Language is: ", value)
	}

}
