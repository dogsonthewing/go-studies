package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 1
const delay = 1

func main() {
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
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

	//declarando um slice
	// sites := []string{"https://www.netflix.com/browse", "https://www.google.com", "https://www.youtube.com", "https://mail.google.com/mail/u/0/"}

	sites := leSitesDoArquivo()

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
	resp, err := http.Get(site)

	errorChecker(err)

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "está online")
		registraLog(site, true)
	} else {
		fmt.Println("O site:", site, "está com problema. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arq, err := os.Open("sites.txt")
	errorChecker(err)

	leitor := bufio.NewReader(arq)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arq.Close()
	return sites
}

func errorChecker(err error) {
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
}

func registraLog(site string, status bool) {
	arq, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	errorChecker(err)
	arq.WriteString(time.Now().Format("2006-01-02 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arq.Close()
}

func imprimeLogs() {
	//ioutil abre e fecha o arquivo dispensando o arq.close()
	fmt.Println("Exibindo logs")
	arq, err := ioutil.ReadFile("log.txt")

	errorChecker(err)
	fmt.Println(string(arq))
}
