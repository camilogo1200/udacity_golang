package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	fmt.Printf("start - time.Now() =>  %v\n", start)

	finish := time.Now()
	fmt.Printf("finish - time.Now() =>  %v\n", start)
	//elapsed time
	elapsedTime := finish.Sub(start)
	fmt.Printf("elapsed time =>  %v\n", elapsedTime)
	addedDate := start.AddDate(0, 0, 2)
	fmt.Printf("added 2 days =>  %q\n", addedDate)
	hour := start.Hour()
	minutes := start.Minute()
	seconds := start.Second()
	day := start.Day()
	month := start.Month()
	year := start.Year()
	fmt.Printf("year = %d, month = %d, day = %d Hour = %d, Minutes = %d, seconds = %d\n", year, month, day, hour, minutes, seconds)
	formattedRfc822 := start.Format(time.RFC822)
	fmt.Printf("formatted RFC-822 => %v\n", formattedRfc822)
	formattedRFC822Z := start.Format(time.RFC822Z)
	fmt.Printf("formatted RFC822Z => %v\n", formattedRFC822Z)
	formattedRuby := start.Format(time.RubyDate)
	fmt.Printf("formatted RubyDate => %v\n", formattedRuby)
	formattedStamp := start.Format(time.Stamp)
	fmt.Printf("formatted Stamp => %v\n", formattedStamp)
	formattedDateTime := start.Format(time.DateTime)
	fmt.Printf("formatted DateTime => %v\n", formattedDateTime)
	formattedDateOnly := start.Format(time.DateOnly)
	fmt.Printf("formatted DateOnly    => %v\n", formattedDateOnly)
	formattedTimeOnly := start.Format(time.TimeOnly)
	fmt.Printf("formatted TimeOnly    => %v\n", formattedTimeOnly)
	formattedRFC1123Z := start.Format(time.RFC1123Z)
	fmt.Printf("formatted RFC1123Z => %v\n", formattedRFC1123Z)
	//ISO week ranges freom 1 to 52 or 53
	yearISO, weekISO := start.ISOWeek()
	fmt.Printf("ISO Year => %v, Week => %v\n", yearISO, weekISO)

	location := start.Location()
	fmt.Printf("location => %v\n", *location)

	nameZone, offset := start.Zone()
	fmt.Printf("Zone name => %s, offset=> %v\n", nameZone, offset)

	unixTime := start.UnixMilli()
	fmt.Printf("Unix time %d\n", unixTime)
	newTicker := time.NewTicker(1 * time.Second)
	defer newTicker.Stop()
	done := make(chan bool)
	go sleepT(done)

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-newTicker.C:
			fmt.Printf("Current time => %v\n", t)
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}

	}
}

func sleepT(done chan bool) {
	time.Sleep(100 * time.Second)
	done <- true
}
