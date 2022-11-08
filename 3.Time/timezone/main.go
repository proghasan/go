package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Dhaka")
	time.Local = loc
	timeFormat := time.Now().Format("02-01-2006 15:04:05 3:04 PM")
	format2 := time.Now().Format("02-01-2006 3:04 PM")

	fmt.Println("DD-MM-YYYY H:m:s H:m AM/PM", timeFormat)
	fmt.Println("DD-MM-YYYY H:m AM/PM", format2)

}
