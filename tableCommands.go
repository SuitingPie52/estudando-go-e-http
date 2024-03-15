package main

import (
	"bufio"
	"fmt"
	"os"
)

var TableCommands = map[string]Explainable{

	"git init": GitInit{},
}

func isCommandValid(c Explainable) bool {

	return c != nil

}

func input() string {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	return s

}

func Run() {

	scanner := bufio.NewScanner(os.Stdin)
	var searched []Explainable
	var userInput, exit string

	for exit != "s" {

		fmt.Printf("Você quer aprender sobre qual comando? ")
		userInput = input()

		if !isCommandValid(TableCommands[userInput]) {

			fmt.Printf("Comando não existe, digite novamente.\n")
			continue

		}

		c := TableCommands[userInput].ConstructObject()

		fmt.Printf(c.explanation)

		searched = append(searched, c)

		fmt.Printf("Você deseja finalizar o programa e imprimir todas as funções que você pesquisou? (s/n) ")
		scanner.Scan()
		exit = scanner.Text()

	}

}
