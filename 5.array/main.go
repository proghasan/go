package main

import "fmt"

func main() {
	fmt.Println("Welcome to array lesson")
	fmt.Println("---------------------------")
	fmt.Println()
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "potato"

	fmt.Println("Fruit List", fruitList, "and length", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Veg list", vegList, "and length", len(vegList))

	fmt.Printf("2D array\n===========\n")
	var courses = [2][3]string{{"GO", "PHP", "JAVA"}, {"MySQL", "NOSQL", "ORACLE"}}
	fmt.Println("Courses list", courses, "and length", len(courses))
}
