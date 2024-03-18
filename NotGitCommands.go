package main

type ChangeDirectory struct {
	Command
}

func (c ChangeDirectory) ConstructObject() Command {

	c.git = false
	c.explanation = "Troca a pasta que você está alterando.\n\n"
	return c.Command

}

type MakeDirectory struct {
	Command
}

func (m MakeDirectory) ConstructObject() Command {

	m.git = false
	m.explanation = "Cria uma nova pasta.\n\n"
	return m.Command

}
