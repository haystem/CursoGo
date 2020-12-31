package main

import("fmt")

func main(){
  channel := make(chan int,2)
  channel <- 45;
  channel <- 44;
  fmt.Println(<-channel)
  fmt.Println(<-channel)
}