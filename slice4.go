package main 

import("fmt")

func main(){
  novoSlice := [][]int {
      []int{1,2,3},
      []int{2,3,4},
      []int{3,4,5},
  }
  fmt.Println(novoSlice)
  fmt.Println("")
  
  for i:=0; i < len(novoSlice); i++{
    fmt.Printf("%d \t \n",novoSlice[0][i])
    for j:=0; j< len(novoSlice);j++{
      fmt.Printf("%d \t \n",novoSlice[j][0])
    }
  }
}