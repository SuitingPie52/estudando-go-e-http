package main 

import "fmt"

type GitInit struct {

	Command

}

func (g GitInit) PrintInfo () {

	fmt.Printf("Transforma a pasta atual em uma pasta git.")

}
