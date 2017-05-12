package main

import "fmt"

func vals() (int, int) {
  return 3, 7
}

func coisa(a, b int) (c , d int) {
  c = a * 2
  d = b * 3
  return c, d
}

func main() {
  a, b := vals()
  fmt.Println("a:",a)
  fmt.Println("b:",b)

  _, c := vals()
  fmt.Println("c:",c)

  d, e := coisa(2, 3)
  fmt.Println("d:",d)
  fmt.Println("e:",e)
}
