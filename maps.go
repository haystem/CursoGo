package main

import ("fmt")

func main(){
  vmap := map[string]int{
      "ana":9898498,
      "alberto":98494849,
      "aline":98494058,
      "cintia":9884945,
  }
  fmt.Println(vmap["alberto"])
  
  if vmap, ok := vmap["alberto"]; !ok{
    fmt.Println("nçao tem",vmap)
  }else{
    fmt.Println("possui",vmap)
  }
}