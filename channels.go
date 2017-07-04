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

func worker2(done chan bool) {
	fmt.Print("ready to work...")
	time.Sleep(time.Second * 3)
	fmt.Println("Work done!")
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

	done := make(chan bool, 2)
	go worker2(done)
	go worker(done)

	<-done
	<-done
	fmt.Println("Acabou")

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	go ping(pings, "Digui, digui diói")
	fmt.Println("fim do mundo")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg + ", digui diói, popoi..."
}
