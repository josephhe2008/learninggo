package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 *time.Second)

	// Timers represent a single event in the future. 
	// You tell the timer how long you want to wait, 
	// and it provides a channel that will be notified
	// at that time. This timer will wait 2 seconds.
	<-timer1.C
	fmt.Println("Timer1 fired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()
	time.Sleep( 5 * time.Second)
	stop2 := timer2.Stop()
	fmt.Println(stop2)
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}