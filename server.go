package main

import (
	"log"
	"net/http"
	"time"
)

type MainHandler struct {
}

func (h MainHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	res.Write([]byte("-----------------------------------\n" +
		"|      DICIONÁRIO DE TERMINAL     |\n" +
		"-----------------------------------\n" +
		"| 1 - Aprender sobre um comando   |\n" +
		"| 2 - Listar todos os comandos    |\n" +
		"-----------------------------------\n"))

	option := req.URL.Query().Get("option")

	if option == "" {

		return

	}

	if !isMenuOptionValid(option) {

		res.Write([]byte("Essa opção não existe no menu"))
		return

	}

	TableServerMenu[option](res, req)

}

func CreateServer() {

	s := &http.Server{
		Addr:         "172.22.88.142:8080",
		Handler:      MainHandler{},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())

}
