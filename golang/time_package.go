package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	fmt.Println(currTime)
	fmt.Printf("%T\n", currTime)

	//format date dd-mm-yyyy
	//golang released on 02-01-2006 so this to format
	// Monday capital M for current day
	// current time
	formatted := currTime.Format("02-01-2006, Monday,15:04:05")
	fmt.Println("format :", formatted)
	//format : 04-04-2025, Friday,16:51:50

	formatted = currTime.Format("2006/01/02, 3:04 PM")
	fmt.Println("format :", formatted)
	//format : 2025/04/04, 4:51 PM

	layout_str := "02/01/2006"

	dateStr := "25/11/2030"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formatted time:", formatted_time)

	// add 1 more day in currentTime
	new_date := currTime.Add(24 * time.Hour)
	fmt.Println("new_date time: ", new_date)
	formatted_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println("formatted_new_date time: ", formatted_new_date)
}
