package main

import(
	"net/http"
	"encoding/json"
	"errors"
	"strconv"
)

var CommandController = map[string]map[string]func(http.ResponseWriter, *http.Request)error {

	"POST": PostFunctions,
	"GET": GetFunctions,
}

var PostFunctions = map[string]func(http.ResponseWriter, *http.Request)error {

	"/command": Create,

}

var GetFunctions = map[string]func(http.ResponseWriter, *http.Request)error {
	
	"/":        func(http.ResponseWriter, *http.Request)error{return nil},
	"/command": SearchByInput,
	"/command/list": ListAllCommands,
	
}

func Create (res http.ResponseWriter, req *http.Request) error {

	var c Command

	/*err :=*/ json.NewDecoder(req.Body).Decode(&c)


	/*
	
	if err != nil {
		
		m, _ := json.Marshal(HTTPMessages[417])
		res.Write(m) 
		
		return err	
	
	}
	
	*/
	
	TableCommands[c.Input] = c
	
	m, _ := json.Marshal(HTTPMessages[201])
	res.Write(m)
	
	return nil

}

func SearchByInput (res http.ResponseWriter, req *http.Request) error {

	if (TableCommands[req.URL.Query().Get("Input")] == Command{}) {
	
		m, _ := json.Marshal(HTTPMessages[404])
		res.Write(m) 
		
		return errors.New("Não existe esse comando.")
	
	}

	c, _ := json.Marshal(TableCommands[req.URL.Query().Get("Input")])
	
	res.Write(c)
	return nil

}

func ListAllCommands(res http.ResponseWriter, req *http.Request)error {
	
	ListCommands := map[string][]Command{}

	for _, c := range TableCommands {

		ListCommands[strconv.FormatBool(c.Git)] = append(ListCommands[strconv.FormatBool(c.Git)], c)

	}
		
	l, _ := json.Marshal(ListCommands)	
	res.Write(l)
	
	return nil

}
