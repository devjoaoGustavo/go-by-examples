package main

import "fmt"

func zeroval(ival int) {
  ival = 0
}

func zeroptr(iptr *int) {
  *iptr = 1000
}

func main() {
  i := 6
  fmt.Println("initial:", i)

  zeroval(i)
  fmt.Println("zeroval:", i)

  zeroptr(&i)
  fmt.Println("zeroptr:", i)

  fmt.Println("pointer:", &i)
}
