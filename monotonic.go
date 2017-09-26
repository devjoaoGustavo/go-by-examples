package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	x := time.NewTimer(time.Second)
	<-x.C
	t := time.Now()
	fmt.Println()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
