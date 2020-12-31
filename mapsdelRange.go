package main

import("fmt")

var total int

func main(){
  
  //map tem key , value diferente de slice que tem indice
  
  vmap := map[int]string{
    11:"Muito Legal",
    22:"Muito Legal em dobro",
    23:"É o muito legal mais um",
    24:"Hummm.. Número da sorte",
    45:"Piri Para ra Pururu",
  }
  
  for key, value := range vmap{
    fmt.Printf("%d, %s \n", key,value)
  
    total+= key
  }
  fmt.Println(total)
  
  delete(vmap, 45)
  fmt.Println(total)
  fmt.Println(vmap)
}