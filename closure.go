package main

import("fmt")

func main(){
 
  a:= soma()(3)
  fmt.Println(a)
}

func soma() func(int) int {
   v:=0
 return func(x int) int{
   v+=x
  return v
 }
}