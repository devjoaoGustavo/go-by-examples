package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1432
	m["b"] = 554
	fmt.Println("map1:", m)

	v1, bla := m["a"]
	fmt.Println("get with second value:", v1, bla)

	fmt.Println("len:", len(m))

	delete(m, "b")
	fmt.Println("map after deleting:", m)

	z_val, prs := m["b"]
	fmt.Println("prs:", prs, z_val)

	_, coisa := m["k2"]
	fmt.Println("the key was present:", coisa)

	n := map[string]int{"age": 15, "point": 25, "timestamp": 1493737290}
	fmt.Println("map:", n)
}
