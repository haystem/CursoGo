package main

import ("fmt")

func main(){
  
  adc := make([]string,10,10)
  
  
  adc = []string{"Ana","Carla","Souza"}
  adc = append(adc,"Ana","Carla","Souza")
  fmt.Println(adc)
}

