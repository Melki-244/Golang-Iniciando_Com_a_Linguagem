package main

import (
  "fmt"
  "os"
  "net/http"
  "time"
)

const delayMonitoramento = 5
const testesMonitoramento = 3  

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

  fmt.Println("")
  
  return comandoLido 
}


func iniciarMonitoramento()  {
  fmt.Println("Monitorando...")  
  
  site := leSitesDoArquivo()
  for atual:= 0; atual < testesMonitoramento; atual++ {
    for i, site := range site {
      fmt.Println("Site:", i ,":", site)
      testeSite(site)
    }
    seconds := time.Second
    time.Sleep(delayMonitoramento * seconds)
    fmt.Println("")
  }
  
  fmt.Println("")

}
func testeSite(site string)  {
  resp, err := http.Get(site)
 
  if err != nil {
    fmt.Println("Ocorreu Um Erro", err) 
  }
  if resp.StatusCode == 200 {
    fmt.Println("Site", site, "Carregado Com Sucesso!") 
  }else{
    fmt.Println("Site", site, "Não Foi Carregado Com Sucesso. Status Code:",resp.StatusCode)
  }
}
func leSitesDoArquivo() []string {
  var sites []string 

  arquivo, _ := os.Open("sites.txt")  
  fmt.Println(arquivo)

  return sites
}
