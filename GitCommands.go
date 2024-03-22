package main

import "net/http"

type GitInit struct {
	Command
}

func (g GitInit) ConstructObject() Command {

	g.explanation = "Transforma a pasta atual em uma pasta git.\n\n"
	g.git = true
	return g.Command

}

func (g GitInit) PrintExplanationInServer(res http.ResponseWriter) {

	res.Write([]byte(g.ConstructObject().explanation))

}

type GitStatus struct {
	Command
}

func (g GitStatus) ConstructObject() Command {

	g.explanation = "Verifica se a pasta git está em dia com o repositório na nuvem.\n\n"
	g.git = true
	return g.Command

}
