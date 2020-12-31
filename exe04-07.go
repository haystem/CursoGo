package main

import("fmt")

func main(){
  pessoa := [][]string{
    []string{"Alberto","Rocha","Pintar"},
    []string{"Joana","Carlos" ,"Decorar"},
    []string{"Mia","Miranda" ,"Levitar"},
  }
  for _,val:= range pessoa{
      fmt.Println(val[0])
      for ind,valor := range val{
        fmt.Println("\t",ind,valor)
      }
  }
}