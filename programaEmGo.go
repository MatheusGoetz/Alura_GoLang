package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Desculpe, não reconheco este comando.")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	var nome string
	versao := 1.1

	fmt.Println("Digite seu nome:")
	fmt.Scan(&nome)

	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa esta na versao", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- sair")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento...")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for i, site := range sites {
		fmt.Println("Testando site", i, ":", site)
		testandoSites(site)
	}
}

func testandoSites(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "carregou com sucesso!")
	} else {
		fmt.Println("Site:", site, "Esta com problemas, status code:", resp.StatusCode)
	}
}
