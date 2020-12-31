package main

import("fmt")

func main(){
  x:=11
  fmt.Println(x)
      //valor(x)
      ponteiro(&x)
       fmt.Println(x)
}

func valor(x int){
  x++
  fmt.Println("Na função", x)
}

func ponteiro(x *int) {
  *x++  
 fmt.Println("Na função", *x)
 fmt.Println("Na função", &x)
 fmt.Println("Na função", *x -1)
 fmt.Println("Na função", &x -1)
}