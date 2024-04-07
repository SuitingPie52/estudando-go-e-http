package main

/*

import (
	"bufio"
	"fmt"
	"os"
)



var TableMenu = map[string]func(){

	"1":  explainCommand,
	"2":  ListAllCommands,
	"-1": func() { os.Exit(0) },
}

func isMenuOptionValid(s string) bool {

	return TableMenu[s] != nil

}

func ListAllCommandsInConsole() {

	ListCommands := map[bool][]string{}

	for name, c := range TableCommands {

		ListCommands[c.ConstructObject().git] = append(ListCommands[c.ConstructObject().git], name)

	}

	fmt.Printf("Not Git:\n\n")

	for _, c := range ListCommands[false] {

		fmt.Println(c)

	}

	fmt.Printf("\nGit:\n\n")

	for _, c := range ListCommands[true] {

		fmt.Println(c)

	}

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
		"| 2 - Listar todos os comandos    |\n" +
		"| -1 - Sair                       |\n" +
		"-----------------------------------\n")

}

func Run() {

	var option string

	for {

		printMenu()

		option = input()

		if !isMenuOptionValid(option) {

			fmt.Println("Essa opção não existe. Escolha outra.")
			continue

		}

		TableMenu[option]()

	}

}

*/
