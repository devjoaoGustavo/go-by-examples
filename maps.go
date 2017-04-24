package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1432
	m["b"] = 554
	fmt.Println("map1:", m)

	v1, bla := m["a"]
	fmt.Println("getwith second value:", v1, bla)

	fmt.Println("len:", len(m))

	delete(m, "b")
	fmt.Println("map after deleting:", m)

	val, prs := m["b"]
	fmt.Println("prs:", prs, val)

	n := map[string]int{"age": 15, "point": 25}
	fmt.Println("map:", n)
}
