package main

import (
	"fmt"
	"time"
)

/*

✅ Golden Tip for Programmers

Whenever you use time, always think:
Format → Format() & Parse()
Arithmetic → Add(), Sub(), Since()
Control → Sleep(), Timer, Ticker
Zone → UTC(), In(), LoadLocation()

Why this weird date?

It’s a memorable sequence:
Year → 2006
Month → 01
Day → 02
Hour → 15 (3PM, 24h format)
Minute → 04
Second → 05
Together: 2006-01-02 15:04:05

Easy to remember as: "1 2 3 4 5 6" (sort of sequential).
*/

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// now to formate
	fmt.Println("\nNow To Formate")
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.Format("Jan 02, 2006 03:04:05 PM"))

	// parse date
	fmt.Println("\nParse Date")
	date, _ := time.Parse("2006-01-02 15:04:05", "2024-05-02 00:10:10")
	fmt.Println(date)

	// time zone
	fmt.Println("\nTime Zone")
	loc, _ := time.LoadLocation("Asia/Kolkata")
	fmt.Println(time.Now().In(loc))

	// utc
	fmt.Println(time.Now().UTC())

	// duration and arithmetic
	fmt.Println("\nDuration and Arithmetic")
	start := time.Now()
	time.Sleep(2 * time.Second)
	fmt.Println("End: ", time.Since(start))

	fmt.Println("Add: ", start.Add(12*time.Hour))

	// timer
	fmt.Println("\nTImer")
	timer := time.NewTimer(3 * time.Second)
	<-timer.C
	fmt.Println("Timer fired")

	// Ticker
	fmt.Println("\nTicker")
	ticker := time.NewTicker(2 * time.Second)
	for i := 0; i < 3; i++ {
		fmt.Println(<-ticker.C)
	}
	ticker.Stop()

	// comparisons
	fmt.Println("\nComparisons")
	t1 := time.Now()
	t2 := t1.Add(10 * time.Minute)
	fmt.Println(t1.After(t2))
	fmt.Println(t1.Before(t2))
	fmt.Println(t1.Equal(t2))

	// sleeping
	fmt.Println("\nSleeping")
	fmt.Println("Strat")
	time.Sleep(2 * time.Second)
	fmt.Print("End")

	// select with time (deadline)
	fmt.Println("\nSelect with time (Deadline)")
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Print("Data: ", res)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}
}
