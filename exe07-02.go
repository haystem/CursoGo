package main

import("fmt")

type pessoa struct{
  nome string
  sobrenome string
  idade int
}

func main(){
  pess := pessoa{"alberto","rocha",29}
  fmt.Println(pess)
  mudeME(&pess)
  fmt.Println(pess)
}

func mudeME(p *pessoa){
  (*p).nome = "Luiz"
  (*p).sobrenome = "Carlicia"
  (*p).idade = 20
}