package main

import "fmt"

type Animal interface {
  Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
  return "Au"
}

type Cat struct {
}

func (c *Cat) Speak() string {
  return "Miau"
}

type Llama struct {
}

func (l Llama) Speak() string {
  return "What should I do?"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
  return "Design Patterns!"
}

func DoSomething(thing interface{}) {
  fmt.Println(thing)
}

func PrintAll(vals []interface{}) {
  for _, val := range vals {
    fmt.Println(val)
  }
}

func main() {
  animals := []Animal{new(Dog), new(Cat), Llama{}, JavaProgrammer{}}
  for _, animal := range animals {
    fmt.Println(animal.Speak())
  }

  DoSomething(Dog{})
  DoSomething(Cat{})
  DoSomething(Llama{})
  DoSomething(JavaProgrammer{})

  coisas := []string{"oi", "mãe", "como", "vc", "está", "?"}
  trens := make([]interface{}, len(coisas))
  for i, coisa := range coisas {
    trens[i] = coisa
  }

  PrintAll(trens)

  fmt.Println(Dog{}.Speak())

  fmt.Println(Cat{}.Speak())
}
