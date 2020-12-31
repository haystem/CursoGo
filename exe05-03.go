package main

import("fmt")

type veiculos struct{
  portas int
  cor string
}

type caminhonete struct{
  veiculos
  quatroRodas bool
}

type sedan struct{
  veiculos
  modeloLuxo bool
}

func main(){
  caminhonete1:= caminhonete{
    veiculos : veiculos {
      portas :2,
      cor : "Vermelho",
    },
    quatroRodas :false,
  }
  
  sedan1:= sedan{
    veiculos : veiculos{
      portas : 4,
      cor : "cinza",
    },
    modeloLuxo : true,
  }
  
  fmt.Println("Caminhonete",caminhonete1,"\n \t",caminhonete1.quatroRodas)
  fmt.Println("sedan",sedan1,"\n \t", sedan1.modeloLuxo)
}