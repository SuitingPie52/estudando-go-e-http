package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type MainHandler struct {
}

func (h MainHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	/*
	
	QUERIA FAZER O MENU SER TRANSFORMADO EM JSON PELO MAPA HTTP MESSAGES, MAS NAO CONSIGO
	DAR '\n' POR JSON
	
	m, _ := json.Marshal(HTTPMessages[1])
	res.Write(m)
	
	*/
	
	res.Write([]byte("--------------------------------------------\n" +
			 "|         DICIONÁRIO DE TERMINAL     	     |\n" +
			 "--------------------------------------------\n" +
			 "| GET:                            	     |\n" +
			 "|                                 	     |\n" +
		 	 "| /: apenas mostra o menu           	     |\n" +
		 	 "| /command: ?Input para achar comando      |\n" +                      
			 "| /command/list: lista todos os comandos   |\n" +
			 "|                                 	     |\n" +
			 "| POST:                                    |\n" +
			 "|                                 	     |\n" +	
			 "| /command: Cadastra novo comando          |\n" +		 
			 "--------------------------------------------\n\n"))
	
	if (!IsMethodAndPathValid(res, req)) {
	
		m, _ := json.Marshal(HTTPMessages[412])
		res.Write(m) 
		
		return
 	
	}
	
	CommandController[req.Method][req.URL.Path](res, req)
	
}

func CreateServer() {

	s := &http.Server{
		Addr:         "192.168.0.9:8080",
		Handler:      MainHandler{},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())

}

func IsMethodAndPathValid (res http.ResponseWriter, req *http.Request)bool {


	return CommandController[req.Method][req.URL.Path] != nil


} 
