package app

import (
	"log"
	"net/http"
)

func Start() {
	// Registering routes/endpoints to DefaultServe Mux- All URLs will be handled by this function
	// Param: 1st: pattern for route or endpoint; 2nd: handler func response we want to send to client
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customersJson", getAllCustomersJsonType)
	http.HandleFunc("/customersXml", getAllCustomersXmlType)

	// starting servers - Create a server listening on port 8000
	// Param: 1st: the port details to start server, 2nd: NIl because we are using default mux
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
