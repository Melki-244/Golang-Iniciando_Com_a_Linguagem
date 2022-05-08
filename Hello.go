package main

import "fmt"
import "os"
import "net/http"

func main()  {
   introducao()   

   fmt.Println("1- Iniciar Monitoramento")
   fmt.Println("2- Exibir os Logs")
   fmt.Println("0- Sair do Programa")
   
   comando := leComando()

   switch comando {
     case 1:
       iniciarMonitoramento()
     case 2:
       fmt.Println("Exibindo Logs")   
     case 0:
       fmt.Println("Saindo do Programa") 
       os.Exit(0) 
     default:
       fmt.Println("Não Conheço Esse Programa")
       os.Exit(-1)
   }
}
func introducao()  {
   nome := "Melki"
   versao := 1.1
   fmt.Println("Olá sr.", nome) 
   fmt.Println("Este Programa Está Na Versão", versao)
}
func leComando() int {
   var comandoLido int
   fmt.Scan(&comandoLido)
   fmt.Println("O comando Escolhido Foi", comandoLido)
   
   return comandoLido
}
func iniciarMonitoramento()  {
  fmt.Println("Monitorando...")  
  site := "https://alura.com.br" 
  http.Get(site)
}











