package main  

import ("fmt")

const (x = iota +1
     y
     z
     t
  )
func main(){
  fmt.Println(x,y,z,t)
}