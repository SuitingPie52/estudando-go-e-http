package main

var TableCommands = map[string]Command{
	"mkdir":      {"mkdir", false, "Cria uma nova pasta."},
	"cd":         {"cd", false, "Troca de pasta"},
	"git init":   {"git init", true, "Transforma a pasta atual em uma pasta Git"},
	"git status": {"git status", true, "Verifica se a pasta remota está em dia com a da nuvem."},
}
