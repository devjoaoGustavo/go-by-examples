package main

import "fmt"

func main() {
	msgs := make(chan string)

	go func() { msgs <- "ping" }()

	go func(w string) {
		p := <-msgs

		fmt.Println(p)

		msgs <- w
	}("Qualcosa")

	msg := <-msgs
	fmt.Println(msg)
}
