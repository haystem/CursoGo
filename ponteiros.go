package main 

import("fmt")

func main(){
  x:= 10
  y:= &x
  
  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(*y) //tenho o endereço e quero saber o que tem dentro dele
  fmt.Println(*&y) //conteudo do endereço
  fmt.Printf("%T, %T", x,y)
}