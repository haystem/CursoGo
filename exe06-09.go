package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	u, s := cfinal(soma, "ola", s1...)
	fmt.Println(u, s)
}

func soma(x ...int) int {
	soma := 0
	for _, valor := range x {
		soma += valor
	}
	return soma
}

func cfinal(f func(x ...int) int, s string, y ...int) (int, string) {
	x := []int{}

	texto := fmt.Sprintf("O valo do texto: %v", s)
	for _, valor := range y {
		x = append(x, valor)
	}
	fu := f(x...)

	return fu, texto
}
