package main

import("fmt")

func main(){
  array := [5]int{1,2,3,4,5}
  fmt.Println(array)
  slice := []int{1,2,3,4,5}
  fmt.Println(slice)
  
  slice2 := append(slice, 7,8,9)
  fmt.Println(slice2)
  fmt.Println(slice[4])
  slice[4] = 23423
  fmt.Println(slice[4])
}