package main

type GitInit struct {
	Command
}

func (g GitInit) ConstructObject() GitInit {

	g.explanation = "Transforma a pasta atual em uma pasta git.\n"
	g.git = true
	return g

}
