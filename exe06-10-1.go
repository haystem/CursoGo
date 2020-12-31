package main 

import("fmt")

func main (){
  
  a := 10
  b := 20
  
  fmt.Println(contar(a))
  fmt.Println(contar(a))
  fmt.Println(contar(a))
  fmt.Println(contar(b))
  fmt.Println(contar(b))
  fmt.Println(contar(b))
}

func contar(x int) func()[]int {
   slice := []int{}
  for y:=0; y<=x; y++ {
     slice = append(slice, y+1)
  } 
  return func() []int {
    for x := range slice {
       s:=[]int
       s:=append(s, x)
    }
  }
}