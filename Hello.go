package main

import "fmt"
import "reflect"

func main()  {
   nome := "Melki"
   idade := 20
   versao := 1.1
   fmt.Println("Olá sr.", nome, "Sua Idade é", idade) 
   fmt.Println("Este Programa Está Na Versão", versao)
   
  fmt.Println("A Variável nome é do Tipo:", reflect.TypeOf(nome))
}
