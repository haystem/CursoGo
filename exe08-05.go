package main

import (
	"fmt"
	"sort"
	"runtime"
	"time"
//	"sync"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byIdade [] user

func (a byIdade) Len() int           { return len(a) }
func (a byIdade) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byIdade) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byLast [] user

func (a byLast) Len() int           { return len(a) }
func (a byLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLast) Less(i, j int) bool { return a[i].Last < a[j].Last }


func main() {
  
  fmt.Println(runtime.NumCPU())
  fmt.Println(runtime.NumGoroutine())
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}
	

	users := []user{u1, u2, u3}

  for _, v := range users{
    sort.Strings(v.Sayings)
  }
	
	go sort.Sort(byLast(users))
	 fmt.Println("Ordenador por sobrenome")
	 ord(users)
	
  fmt.Println("\n")
  go sort.Sort(byIdade(users))
  time.Sleep(30)
 	fmt.Println("Ordenador por idade")
	ord(users)
	
	
  fmt.Println(runtime.NumCPU())
  fmt.Println(runtime.NumGoroutine())
}

func ord(x []user) {
  for i,v := range x{
     fmt.Println("Registro: ",i, "\t Nome: ", v.First,"\t Sobrenome: ", v.Last,"\t Idade: ", v.Age)
     for _, c := range v.Sayings{
       fmt.Println("\t",c)
     }
  }
}