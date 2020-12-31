package main

import ("fmt")

const (x int =10
      a string = "Say my name"
      b float64 = 283.847
      y = true
      t = "Aloha"
      v = 34 
      )
      
      
func main(){

  fmt.Printf("O valor de x: %v tipo: %T \n",x,x)
  fmt.Printf("O valor de y: %v tipo: %T \n",y,y)
  fmt.Printf("O valor de a: %v tipo: %T \n",a,a)
  fmt.Printf("O valor de b: %v tipo: %T \n",b,b)
  fmt.Printf("O valor de t: %v tipo: %T \n",t,t)
  fmt.Printf("O valor de v: %v tipo: %T \n",v,v)
}