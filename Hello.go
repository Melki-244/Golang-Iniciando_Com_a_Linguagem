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

   if comando == 1 {
     fmt.Println("Iniciando Monitoramento") 
   }else if comando == 2 {
     fmt.Println("Exibindo Logs")
   }else if comando == 0 {
     fmt.Println("Saindo Do Programa")
   }else{
     fmt.Println("Não Conheço Seu Comando")
   }
}
