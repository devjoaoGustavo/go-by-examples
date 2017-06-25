package main

import (
	"fmt"
	"os"
)

type SyntaxError struct {
	msg    string
	Offset int64
}

func (e *SyntaxError) Error() string { return e.msg }

type NegativeSqrtError float64

func (f NegativeSqrtError) Error() string {
	return fmt.Sprintf("Math: square root of negative number: %g", float64(f))
}

func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("Math: square root of negative value: %g", n)
	}
	return n * n, nil
}

func main() {
	fmt.Println("Mensagem de erro ao tentar abrir arquivo que não funciona")
	_, err := os.Open("filename.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Outro exemplo")
	_, erro := Sqrt(-1)
	if erro != nil {
		fmt.Println(erro)
	}
}
