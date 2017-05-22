package main

import "fmt"

func intSeq() func() int {
  i := 0
  res := func() int {
    i += 1
    return i
  }

  return res
}

func main() {
  nextInt := intSeq()

  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())

  newInts := intSeq()
  fmt.Println(newInts())
}
