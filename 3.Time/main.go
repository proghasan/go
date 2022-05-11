package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Welcome to time package")
	fmt.Println("Current time", t)
	fmt.Println()

	fmt.Println("Current BD Format")
	// YYY-MM-DD HH:MM:SS
	formatTime := t.Format("02-01-2006 15:04:05 3:04 PM")
	formatTime2 := t.Format("02-01-2006 3:04:05 PM")
	formatWithNanoSecond := t.Format("02-01-2006 3:04:05.000000000")

	fmt.Println("Formatted time", formatTime)
	fmt.Println("Formatted time with am/pm", formatTime2)
	fmt.Println("Formatted with nano time", formatWithNanoSecond)

	// date function
	dateFormat := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Minute(), t.Nanosecond(), time.UTC)
	fmt.Println("Date format with raw", dateFormat)
	fmt.Println("Date format understand human", dateFormat.Format("02-01-2006 3:04:05 PM"))
}
