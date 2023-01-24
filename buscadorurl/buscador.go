package buscadorurl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//go run .\main.go gopl.io

func Buscador() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		f, err := os.Create("arquivos/buscar.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao gerar arquivo do buscar, cause: %v", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)

		f.Write(b)
	}
}
