package main

import "fmt"
import "os"
import "net/http"
import "reflect"

func main()  {
  arrayPadrao()
  arraySlice()
   //introducao()   
  for {
    //exibeMenu() 
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

  site := "https://alura.com.br" 
  resp, _ := http.Get(site)
  
  if resp.StatusCode == 200 {
    fmt.Println("Site", site, "Carregado Com Sucesso!") 
  }else{
    fmt.Println("Site", site, "Não Foi Carregado Com Sucesso. Status Code:",resp.StatusCode)
  }
}
func arraySlice()  {
  nomes := []string {"Melki","Kátia","Kaline","Maria"} 
  fmt.Println("nomes no Slice", nomes)
  fmt.Println(reflect.TypeOf(nomes))
  fmt.Println("O Meu Slice Tem", len(nomes))
  fmt.Println("A Capacidade Do Slice é de", cap(nomes))

  nomes = append(nomes, "Aparecida")

  fmt.Println("nomes no Slice", nomes)
  fmt.Println(reflect.TypeOf(nomes))
  fmt.Println("O Meu Slice Tem", len(nomes))
  fmt.Println("A Capacidade Do Slice é de", cap(nomes))
}
func arrayPadrao()  {
  var sites [4] string
  sites[0] = "https://caelum.com.br" 
  sites[1] = "https://mercadolivre.com.br"
  sites[2] = "https://amazon.com.br"

  fmt.Println("Sites No Array", sites)
  fmt.Println(reflect.TypeOf(sites)) 
}
