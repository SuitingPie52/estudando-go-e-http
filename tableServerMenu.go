package main

import (
	"net/http"
)

var TableServerMenu = map[string]func(http.ResponseWriter, *http.Request){

	"1": ExplainCommandInServer,
	"2": ListAllCommandsInServer,
}

func ExplainCommandInServer(res http.ResponseWriter, req *http.Request) {

	res.Write([]byte("Qual comando você quer que eu explique?\n"))

	c := req.URL.Query().Get("command")

	if c == "" {

		return

	}

	if !isCommandValid(TableCommands[c]) {

		res.Write([]byte("Não existe esse comando."))
		return

	}

	res.Write([]byte(TableCommands[c].ConstructObject().explanation))
	return

}

func ListAllCommandsInServer(res http.ResponseWriter, req *http.Request) {

	ListCommands := map[bool][]string{}

	for name, c := range TableCommands {

		ListCommands[c.ConstructObject().git] = append(ListCommands[c.ConstructObject().git], name)

	}

	res.Write([]byte("Not Git:\n\n"))

	for _, c := range ListCommands[false] {

		res.Write([]byte(c + "\n"))

	}

	res.Write([]byte("\nGit:\n\n"))

	for _, c := range ListCommands[true] {

		res.Write([]byte(c + "\n"))

	}

	return

}
