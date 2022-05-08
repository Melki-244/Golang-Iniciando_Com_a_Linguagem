package main

import (
  "fmt"
  "os"
  "net/http"
)

func main()  {
  introducao()   
  for {
    exibeMenu() 
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
}
func introducao()  {
  nome := "Melki"
  versao := 1.1
  fmt.Println("Olá sr.", nome) 
  fmt.Println("Este Programa Está Na Versão", versao)
}
func exibeMenu()  {
  fmt.Println("1- Iniciar Monitoramento")
  fmt.Println("2- Exibir os Logs")
  fmt.Println("0- Sair do Programa")
}
func leComando() int {
  var comandoLido int
  fmt.Scan(&comandoLido)
  fmt.Println("O Comando Escolhido Foi", comandoLido)

  return comandoLido 
}
func iniciarMonitoramento()  {
  fmt.Println("Monitorando...")  
  
  sites := [] string {"https://alura.com.br", "https://caelum.com.br", "https://amazon.com"}
  for position, content := range sites {
   fmt.Println("Passando A Posição", position,"e o Conteudo", content,"Do Slice") 
  }
  site := "https://alura.com.br" 
  resp, _ := http.Get(site)
  
  if resp.StatusCode == 200 {
    fmt.Println("Site", site, "Carregado Com Sucesso!") 
  }else{
    fmt.Println("Site", site, "Não Foi Carregado Com Sucesso. Status Code:",resp.StatusCode)
  }
}
