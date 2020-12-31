package main

import("fmt")

func main(){
 array := []string{"Ana","Joana","Siraya","Carla","Vivian","Betina","Laritssa","Bianca","Muana","Juressica"} 

  fmt.Println(array[:3])
  fmt.Println(array[4:len(array)])
  fmt.Println(array[1:7])
  fmt.Println(array[2:len(array)-1])
  fmt.Println(array[len(array)-2])
  fmt.Println(len(array))
}
