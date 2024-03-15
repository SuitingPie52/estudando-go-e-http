package main

import (
	"bufio"
	"fmt"
	"os"
)

var TableMenu = map[string]func(){

	"1":  explainCommand,
	"-1": func() { os.Exit(0) },
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

func explainCommand() {

	var userInput string

	for {

		fmt.Printf("Você quer aprender sobre qual comando? ")
		userInput = input()

		if !isCommandValid(TableCommands[userInput]) {

			fmt.Printf("Comando não existe, digite novamente.\n")
			continue

		}

		c := TableCommands[userInput].ConstructObject()

		fmt.Printf(c.explanation)

		return
	}
}

func printMenu() {

	fmt.Printf("-----------------------------------\n" +
		"|      DICIONÁRIO DE TERMINAL     |\n" +
		"-----------------------------------\n" +
		"| 1 - Aprender sobre um comando   |\n" +
		"| -1 - Sair                       |\n" +
		"-----------------------------------\n")

}

func Run() {

	var option string

	for {

		printMenu()

		option = input()
		TableMenu[option]()

	}

}
