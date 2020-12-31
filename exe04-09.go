package main

import("fmt")

func main(){
  vmap := map[string][]string{
     "Alberto Rocha":   []string{"Pintar","Quitar","Feedar"},
     "Joana Maria":     []string{"Softbol","Pinball","Ball"},
     "Dercy Gonçalbes": []string{"Xingar","Rir","Criticar"},
  }

  vmap["Derecy Balabraiano"] = []string{"Quitar","Feeding","Sujar e limpar"}
  
    for chave,valor := range vmap{
    fmt.Println(chave)
    for ind,val:= range valor{
      fmt.Println("\t",ind,val)
    }
  }
  
  
}
