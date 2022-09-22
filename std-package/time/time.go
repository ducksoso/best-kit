package main

import (
	"fmt"
	"time"
)

// CSTLayout China Standard Time Layout
const CSTLayout = "2006-01-02 15:04:05"

func main() {

	countdown()

	RFC3339Str := "2020-11-08T08:18:46+08:00"
	RFC3339ToCSTLayout(RFC3339Str)

	//
	TimeLocation()
}

// countdown  计算一天剩余的时间 - 秒
func countdown() {
	fmt.Println("countdown start...")
	endTime, err := time.Parse(CSTLayout, time.Now().Format("2006-01-02")+" 23:59:59")
	if err != nil {
		panic(err)
	}
	fmt.Println(endTime.String())
	ttlSec := endTime.Sub(time.Now()).Seconds()
	ttlMinutes := endTime.Sub(time.Now()).Minutes()
	fmt.Println(ttlSec, ttlMinutes)
	fmt.Println("countdown end...")

	fmt.Println()
	return
}

// RFC3339ToCSTLayout convert rfc3339 value to China standard time layout
func RFC3339ToCSTLayout(value string) {
	fmt.Println("RFC3339ToCSTLayout start...")
	// 时区
	// time.RFC3339
	var cst *time.Location
	var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}

	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(err)
	}

	format := ts.In(cst).Format(CSTLayout)
	fmt.Println(format)
	fmt.Println("RFC3339ToCSTLayout end...")
	fmt.Println()

}

func TimeLocation() {
	fmt.Println("TimeLocation start...")

	// Pretend this is the date and time we extracted
	year := 2013
	month := 9
	day := 14
	hour := 15
	minute := 6

	// Capture the location based on the timezone id from Google
	// location, err := time.LoadLocation("America/Chicago")
	// location, err := time.LoadLocation("Asia/Shanghai")
	// if err != nil {
	// 	fmt.Printf("ERROR : %s", err)
	// 	return
	// }

	// Capture the local and UTC time based on timezone
	// localTime := time.Date(year, time.Month(month), day, hour, minute, 0, 0, location)
	localTime := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.Local)
	utcTime := localTime.UTC()
	local := localTime.Local()
	fmt.Printf("Local Time: %v\n", localTime)
	fmt.Printf("UTC Time: %v\n", utcTime)
	fmt.Printf("local Time: %v\n", local)

	fmt.Println("TimeLocation end...")

	fmt.Println()

}
