package main

import "fmt"

func main() {

	fmt.Printf("Você quer aprender sobre qual comando?")
	
	TableCommands["git init"].PrintInfo()
	
}	
