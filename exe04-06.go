package main

import("fmt")

func main(){
  estados:= make([]string,15,20)
  estados = []string{
    "Acre", 
    "Alagoas",
    "Amapá", 
    "Amazonas",
    "Bahia", 
    "Ceará", 
    "Espírito Santo", 
    "Goiás", 
    "Maranhão", 
    "Mato Grosso",
    "Mato Grosso do Sul",
  }
  estados = append(estados,"Minas Gerais", "Pará", "Paraíba", 
  "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", 
  "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", 
  "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins" )

  fmt.Println(len(estados),cap(estados))

for x:=0; x < len(estados);x++{
  fmt.Println(estados[x])
}
  
}
