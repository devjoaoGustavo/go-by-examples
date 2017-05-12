package main

import "fmt"

func main() {
  nums := [...]int{2,3,4}
  sum := 0

  for _, num := range nums {
    sum += num
  }
  fmt.Println("Soma:", sum)

  for i, num := range nums {
    if num == 3 {
      fmt.Println("Index:", i)
    }
  }

  kvs := map[string]string{"a": "apple", "b": "banana", "c": "coffee"}
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
  }

  for k := range kvs {
    fmt.Println("Key:", k)
  }

  for _, v := range kvs {
    fmt.Println("Value:", v)
  }

  for i, c := range "golang" {
    fmt.Printf("%d - %c\n", i, c)
  }

  for i, c := range "go" {
    fmt.Println(i, c)
  }
}
