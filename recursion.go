package main

import "fmt"

func fact(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact(n-1)
}

func main() {
  fmt.Println("Fatorial de 7 -", fact(7))
  fmt.Println("Fatorial de 1 -", fact(1))
  fmt.Println("Fatorial de 3 -", fact(3))
  fmt.Println("Fatorial de 10 -", fact(10))
}
