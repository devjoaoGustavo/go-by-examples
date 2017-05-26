package main

import (
  "fmt"
  "math"
)

type geometry interface {
  area() float64
  perim() float64
}

type rect struct {
  width, height float64
}

type circle struct {
  radius float64
}

type coisa struct {
  largura float64
  altura float64
}

func (x coisa) area() float64 {
  return x.largura * x.altura
}

func (r rect) area() float64 {
  return r.width * r.height
}

func (r rect) perim() float64 {
  return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
  return 2 * math.Pi * c.radius
}

func measure(g geometry) {
  fmt.Println(g)
  fmt.Println(g.area())
  fmt.Println(g.perim())
}

func main() {
  r := rect{width: 3, height: 4}
  c := circle{radius: 5}

  measure(r)
  measure(c)
  // a instrução abaixo provocará um erro, poisa estrutura coisa não implementa a interface geometry
  measure(coisa{largura: 10, altura: 2})
}
