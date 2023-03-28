package serverlissajous

import (
	"log"
	"net/http"
	"strconv"

	"github.com/mixcleiton/livro-a-linguagem-de-programa-o-go/lissajous"
)

func NewServer() {
	http.HandleFunc("/", handler) //cada requisição chama handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler ecoa o componente Path do URL requisitado
func handler(w http.ResponseWriter, r *http.Request) {
	cycle, err := strconv.Atoi(r.FormValue("cycle"))
	if err != nil {
		cycle = 5
	}

	lissajous.Lissajous(w, lissajous.ValuesLissajous{
		Cycles:  cycle,
		Res:     0.001,
		Size:    100,
		NFrames: 64,
		Delay:   8,
	})
}
