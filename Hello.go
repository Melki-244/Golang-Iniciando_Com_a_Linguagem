package main

import "fmt"

func main()  {
   nome := "Melki"
   versao := 1.1
   fmt.Println("Olá sr.", nome) 
   fmt.Println("Este Programa Está Na Versão", versao)
   
   fmt.Println("1- Iniciar Monitoramento")
   fmt.Println("2- Exibir os Logs")
   fmt.Println("0- Sair do Programa")
   
   var comando int
   fmt.Scan(&comando)
   fmt.Println("O comando Escolhido Foi", comando)
  
   switch comando {
     case 1:
       fmt.Println("Monitorando...")  
     case 2:
       fmt.Println("Exibindo Logs")   
     case 0:
       fmt.Println("Saindo do Programa") 
     default:
       fmt.Println("Não Conheço Esse Programa")
   }
}










