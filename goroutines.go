package main

import "fmt"

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutines")

	go f("são legais")

	go func(msg string) {
		fmt.Println(msg)
	}("Hello World!!")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
