package main

import (
	"fmt"
	"time"
)

// A small programme written to get some understanding of how time zone handling is done in Go.
func main() {

	// Input, perhaps from user via an HTML form or some other means
	fromStr := "2013-07-09"
	toStr := "2013-11-10" // 24 hours will be added to this date below to make it exclusive (eg. include all day!)
	locationStr := "Europe/Copenhagen"

	// Get location, fall back to UTC
	location, err := time.LoadLocation(locationStr)
	if err != nil {
		location = time.UTC
	}

	// Parse dates, fall back to today for both dates if necessary

	from, err := time.ParseInLocation("2006-01-02", fromStr, location)
	if err != nil {
		fmt.Printf("Unable to parse from-date '%v'\n", fromStr)
		from, _ = time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), location)
	}

	to, err := time.ParseInLocation("2006-01-02", toStr, location)
	if err != nil {
		fmt.Printf("Unable to parse to date '%v'\n", toStr)
		to, _ = time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), location)
	}

	// Want all day of "to" - so add 24 hours and use it as exclusive when doing database queries for example.
	to = to.Add(24 * time.Hour)

	fmt.Printf("%11s %v\n%11s %v\n%11s %v.\n", "Period from", from, "to", to, "in", location)
	fmt.Printf("Notice offset. Daylight saving time means that %v can be either %v or %v.\n", location, from.Format("-0700"), to.Format("-0700"))
	fmt.Printf("In UTC these dates are from '%v' to '%v' which is useful for querying databases that stores dates in UTC.\n", from.UTC(), to.UTC())
}
