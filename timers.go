package main

import (
	"fmt"
	"time"
)

func main() {
	// func NewTimer(d Duration) *Timer
	timer1 := time.NewTimer(time.Second * 2)
	x := <-timer1.C
	fmt.Println("Timer 1 expired", x)
	// func NewTimer(d Duration) *Timer
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	// func (t *Timer) Stop() bool
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
