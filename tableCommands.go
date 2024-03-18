package main

var TableCommands = map[string]Explainable{
	"mkdir":      MakeDirectory{},
	"cd":         ChangeDirectory{},
	"git init":   GitInit{},
	"git status": GitStatus{},
}
