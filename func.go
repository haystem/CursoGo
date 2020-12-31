package main

import("fmt")


type funcionario struct{
  nome string
  sobrenome string
  departamento
  grupo
}

type grupo struct{
   nome string
   cidade string
   estado string
}

type departamento struct{
    setor string
    salario float64
}

func (f funcionario) imprimirFuncionario() {
   fmt.Println("O nome do funcionario: ", f.nome)
}

func (f funcionario) imprimirFuncionario1() {
   fmt.Println("O sobrenome do funcionario: ", f.sobrenome)
   fmt.Println("O funcionario so setor: ", f.setor)
   fmt.Println("O funcionario do grupo : ", f.grupo.nome)
}

func (f funcionario) imprimirGrupo(){
  fmt.Printf("O grupo %v tem salario de %.2f",f.grupo.nome, f.salario)
}

type metodos interface{
  imprimirGrupo()
  imprimirFuncionario1()
  imprimirFuncionario()
}

func prova(m metodos) {
  m.imprimirFuncionario()
  m.imprimirFuncionario1()
  m.imprimirGrupo()
}

func main(){
  
  pessoa1 := funcionario{
    nome: "Alberto",
    sobrenome : "Rocha",
    departamento:departamento{setor:"Cozinha", salario:2500.00},
    grupo: grupo{nome:"Ferreira",cidade:"Rio de Janeiro",estado:"RJ"},
  } 
  
 prova(pessoa1)
  
}