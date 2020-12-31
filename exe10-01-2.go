package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)

	go func() {
		c <- 42
		c <- 43
		close(c)
	}()

	for v := range c {
		fmt.Println(v, <-c)
	}
}
