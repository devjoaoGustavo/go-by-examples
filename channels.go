package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("working...")
	s := time.Second
	time.Sleep(s)
	fmt.Println("done", s)
	done <- true
}

func main() {
	messages := make(chan string)

	go func(msg string) { messages <- msg }("ping")

	msg := <-messages
	fmt.Println(msg)

	commands := make(chan string, 2)

	go func(c chan string) {
		c <- "Hello"
	}(commands)

	go func(c chan string) {
		c <- "World"
	}(commands)

	go func(c <-chan string) {
		fmt.Println(<-c)
	}(commands)
	fmt.Println(<-commands)

	done := make(chan bool, 1)
	go worker(done)

	<-done
	fmt.Println("Acabou")

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	go ping(pings, "Popoi")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	fmt.Println("fim do mundo")
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
