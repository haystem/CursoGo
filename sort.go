package main

import(
  "sort"
  "fmt"
  )
func main(){
  intr:= []int{1,2,4,5,6,7,8,9,0}
  flt:= []float64{1.5,2.23,2.4,3.5,1.6,0.7,2.8,1.9,2.0}
  str:=[]string{"Abacaxi","Pera","Uva","Maça","Abacate"}
  
  sort.Float64s(flt)
  sort.Ints(intr)
  sort.Strings(str)
  
  //[]string = intr
  //[]string = flt
  
///  os.Stdout(flt)
  fmt.Println(flt)
}
