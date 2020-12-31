package main

import (
	"fmt"
)

func main() {

	total, quantos, oi := soma(10, 10, 1, 2, 3, 5)

	fmt.Println(total, quantos, oi)
}

func soma(x ...int) (int, int, string) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), "Bom dia!"
}
