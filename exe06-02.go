package main

import ("fmt")

func main(){
  s1 := []int{1,2,3,4,56,75}
  fmt.Println(soma(s1...))
  fmt.Println(soma2(s1))
  
}

func soma(x ... int) int {
  total := 0
  for _, v :=range x{
   total += v
  }
  return total
}

func soma2(x []int) int{
  total := 0
  for _, v := range x{
    total += v
  }
  return total
}