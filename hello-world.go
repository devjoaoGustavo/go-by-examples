package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Printf("jogando %d\n", i)
		i++
		if i > 100 {
			break
		}
	}
	fmt.Println("Sai do loop")
}
