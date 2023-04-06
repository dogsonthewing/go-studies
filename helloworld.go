package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs")
		case 3:
			fmt.Println("Saindo")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
	}

}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando")
	sites := []string{"https://www.netflix.com/browse", "https://www.google.com", "https://www.youtube.com", "https://mail.google.com/mail/u/0/"}

	//for range
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			reflect.TypeOf(i)
			testaSite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "está online")
	} else {
		fmt.Println("O site:", site, "está com problema. Status code:", resp.StatusCode)
	}
}
