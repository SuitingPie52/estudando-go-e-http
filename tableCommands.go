package main

import "fmt"

var TableCommands = map[string]Explainable{

	"git init": GitInit{},
}

func isCommandValid(c Explainable) bool {

	return c != nil

}

func Run() {

	var searched []Explainable
	var userInput, exit string

	for exit != "s" {

		fmt.Printf("Você quer aprender sobre qual comando? ")
		fmt.Scanln(&userInput)

		c := TableCommands[userInput].ConstructObject()

		if !isCommandValid(c) {

			fmt.Printf("Comando não existe, digite novamente.")
			continue

		}

		fmt.Printf(c.explanation)

		searched = append(searched, c)

		fmt.Printf("Você deseja finalizar o programa e imprimir todas as funções que você pesquisou? (s/n) ")
		fmt.Scanln(exit)

	}

}
