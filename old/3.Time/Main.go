package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()

	// fmt.Println(time.Local().Format("2020-5-10"))
	fmt.Println("MM-DD-YYYY : ", time.Format("01-02-2006"))

}
