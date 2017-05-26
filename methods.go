package main

import "fmt"

type rect struct {
  width, height int
}

func (r rect) area() int {
  return r.width * r.height
}

func (r rect) perim() int {
  return 2*r.width + 2*r.height
}

func (r *rect) SetWidth(newWidth int) {
  r.width = newWidth
}

func (r *rect) SetHeight(newHeight int) {
  r.height = newHeight
}

func main() {
  r := rect{width: 10, height: 5}

  fmt.Println("area:", r.area())
  fmt.Println("perim:", r.perim())

  rp := &r
  fmt.Println("area:", rp.area())
  fmt.Println("perim:", rp.perim())

  fmt.Println("Struct:", r)
  r.SetWidth(15)
  r.SetHeight(10)
  fmt.Println("Modified struct:", r)
  fmt.Println("bigger area:", r.area())
  fmt.Println("bigger perim:", r.perim())
}
