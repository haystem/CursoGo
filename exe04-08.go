package main

import("fmt")

func main(){
   vmap := map[string][]string {
     "Alberto Rocha":   []string{"Pintar","Quitar","Feedar"},
     "Joana Maria":     []string{"Softbol","Pinball","Ball"},
     "Dercy Gonçalbes": []string{"Xingar","Rir","Criticar"},
   }
   for x,y := range vmap{
       fmt.Println(x)
       for index,value := range y{
         fmt.Println("\t",index,value)
       }
 }
 
}