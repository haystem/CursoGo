package main 

import ("fmt")

const (_ = iota
        x = iota + 2020
        y
        t
        b
        )
        
func main(){
   fmt.Println(x,y,t,b)
}