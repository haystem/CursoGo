package main

import("fmt")

func main(){
  ss:= [][]int{  //dois niveis mas pode ter [][][]int
      []int{1,2,3},  // 0 1 2
      []int{3,4,5},  // 1
      []int{9,4,34},// 2
      []int{1,4,5,6,2,3,4},
  }
  fmt.Println(ss[2][2])
  fmt.Println(ss)
}