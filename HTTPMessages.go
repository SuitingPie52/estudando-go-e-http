package main


type HTTPMessage struct {

	Code int `json: "Code"`
	Message string `json: "Message"`

}

var HTTPMessages = map[int]HTTPMessage {
  
  /*
  1: HTTPMessage{Code: 1, Message: ("--------------------------------------------\n" +
			 "|         DICIONÁRIO DE TERMINAL     	     |\n" +
			 "--------------------------------------------\n" +
			 "| GET:                            	     |\n" +
			 "|                                 	     |\n" +
		 	 "| /command: ?Input para achar comando      |\n" +                      
			 "| /command/list: lista todos os comandos   |\n" +
			 "|                                 	     |\n" +
			 "| POST:                                    |\n" +
			 "|                                 	     |\n" +	
			 "| /command: Cadastra novo comando          |\n" +		 
			 "--------------------------------------------\n\n")},
  */
	
	201: {201, "Created"},
	404: {404, "Not Found"},
	412: {412, "Precondition Failed"},
  	417: {417, "Expectation Failed"},
  
}


