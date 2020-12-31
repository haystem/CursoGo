package main

import("fmt")

type pessoa struct{
  nome string
  sobrenome string
  sabores []string
}

func main(){
  adc := make(map[string]pessoa)
  
  adc["JUNIOR"] = pessoa {
    nome: "Alberto",
    sobrenome:"Rocha",
    sabores: []string{"Mamão","Uva","Laranja"}}
    
    adc["Jussiara"] = pessoa {
    nome: "Mermeriza",
    sobrenome:"Diaria",
    sabores: []string{"Kuwieu","Hortelã","Soda"}}
    
   for chave, valor := range adc{
     fmt.Println(chave,"\n",valor.sobrenome,valor.nome)
     fmt.Printf("Os sabores são: \n")
     for _,ind:= range valor.sabores {
        fmt.Println("\t",ind)
     }
   }
}