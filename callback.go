package main

import("fmt")

func main(){
  
  t:= caly(soma,[]int{1,2,3,4,5,6,7,8,9}...)
  fmt.Println(t)
  
}

func soma(x ...int) int{
   soma := 0 
   for _,v := range x{
      soma+=v
   }
   return soma
}


func caly(f func(x ... int) int, y ...int) int{
  var slice []int
   for _,v := range y {
     if v%2 ==0{
       slice = append(slice,v)
     }
   }
   
   total := f(slice ...)
   return total
}