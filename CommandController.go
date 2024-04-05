package main

import(
	"net/http"
	"encoding/json"
)

var CommandController = map[string]map[string]func(http.ResponseWriter, *http.Request)error {

	"POST": PostFunctions,
	"GET": GetFunctions,
}

var PostFunctions = map[string]func(http.ResponseWriter, *http.Request)error {

	"/command": Create,

}

var GetFunctions = map[string]func(http.ResponseWriter, *http.Request)error {

	"/command": SearchByInput,
	
}

func Create (res http.ResponseWriter, req *http.Request) error {

	var c Command

	err := json.NewDecoder(req.Body).Decode(&c)
	
	if err != nil {
	
		return err	
	
	}
	
	TableCommands[c.Input] = c		
	return nil

}

func SearchByInput (res http.ResponseWriter, req *http.Request) error {

	c, _ := json.Marshal(TableCommands[req.URL.Query().Get("Input")])
	res.Write(c)
	return nil

}
