package main

import (
	"fmt"
	"math"
)

const serverType string = "Local Server running..."

func main() {
	fmt.Println(serverType)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(d))

}
