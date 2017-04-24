package main

import "fmt"
import "regexp"
import "io/ioutil"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	fmt.Println("len:", len(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	s = append(s, "d")
	aux := []string{"e", "f"}
	s = append(s, aux...)
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	l = s[(len(s) - 4):]
	fmt.Println("sl4:", l)

	ary := [3]string{"joao", "e", "vale"}
	slc := ary[:]
	fmt.Println("sl5:", slc)

	o := []string{"a", "b", "c", "d", "e"}
	x := o[2:4]
	fmt.Println("len of original:", len(o))
	fmt.Println("len of re-slice:", len(x))
	fmt.Println("cap of original:", cap(o))
	fmt.Println("cap of re-slice:", cap(x))
	fmt.Println("slice:", o)
	fmt.Println("re-slice:", x)

	x[1] = "T"
	fmt.Println("modified slice:", o)

	fmt.Println("re-slice before:", x)
	x = x[:cap(x)]
	fmt.Println("re-slice after:", x)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// copying a slice
	a := []byte{1, 2, 3}
	b := make([]byte, len(a), (cap(a)+1)*2)
	fmt.Println("slices:", a, b)
	fmt.Println("slices caps:", cap(a), cap(b))
	for i := range a {
		b[i] = a[i]
	}
	fmt.Println("slices:", a, b)
	fmt.Println("slices caps:", cap(a), cap(b))
	a = b
	fmt.Println("slices:", a, b)
	fmt.Println("slices caps:", cap(a), cap(b))

	// another way to copy slices
	q := []byte{1, 2, 3}
	w := make([]byte, len(q), (cap(q)+1)*2)
	copy(w, q)
	fmt.Println("slices:", q, w)
	fmt.Println("slices caps:", cap(q), cap(w))
	q = w
	fmt.Println("slices:", q, w)
	fmt.Println("slices caps:", cap(q), cap(w))

	//appending
	y := []byte{1, 2, 3}
	y = AppendByte(y, 4, 5, 6)
	fmt.Println("apd:", y)

	d := make([]int, 5, 7)
	fmt.Println("len:", len(d))
	fmt.Println("cap:", cap(d))

	e := make([]string, 10)
	fmt.Println("len:", len(e))
	fmt.Println("cap:", cap(e))

	f := []int{}
	fmt.Println("len:", len(f))
	fmt.Println("cap:", cap(f))

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + 1
		}
	}
	fmt.Println("2d:", twoD)

	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	final := Filter(lista, IsEven)
	fmt.Println("resultado:", final)

	fmt.Println("find in file:", FindDigits("aux.txt"))

	fmt.Println("find in file copying digits:", CopyDigits("aux.txt"))
}

//appending
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, n)
		copy(newSlice, slice)
		slice = newSlice
	}
	copy(slice[m:n], data)

	return slice
}

func Filter(s []int, fn func(int) bool) []int {
	var p []int
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func FindDigits(filename string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}

func CopyDigits(filename string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
