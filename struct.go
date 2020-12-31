package main

import("fmt")

type x struct{
  nome string
  sobrenome string
  fumante bool
}

func main(){
   cliente1:= x{
     nome:"João",
     sobrenome:"Da Silva",
     fumante: false,
   }
   
   cliente2:=x{
     nome:"Joana",
     sobrenome:"Da Silva",
     fumante : true,
   }
   fmt.Println(cliente2,"\n",cliente1)
}