package main

import "fmt"

func sum(nums ...int) {
  fmt.Print(nums, " ")
  total := 0

  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

func coisa(trem ...interface{}) {
  fmt.Println(trem...)
}

func main() {
  fmt.Println("Dois parâmetros")
  sum(1, 2)
  fmt.Println("Três parâmetros")
  sum(1, 2, 3)

  nums := []int{1, 2, 3, 4}
  nums = append(nums, 5, 6)
  fmt.Println("Quatro parâmetros")
  sum(nums...)

  fmt.Println("varias coisas")
  coisa(1, 2, "alguma", "coisa", nums)
}
