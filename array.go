package main

import("fmt")

var x [5]int // x[5]=9 não irá funcionar pq ele começa do zero e o quinto valor termina número 4

func main(){
  x[0] = 1
  x[1] = 10
  x[3] = 100
  fmt.Println(x[1],x[0],x[3])
  fmt.Println(x)
  fmt.Printf("%T\n",x)
  fmt.Println(len(x))
}