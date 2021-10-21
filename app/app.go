package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// // Registering routes/endpoints to DefaultServe Mux- All URLs will be handled by this function
	// // Param: 1st: pattern for route or endpoint; 2nd: handler func response we want to send to client
	// http.HandleFunc("/greet", greet)
	// http.HandleFunc("/customersJson", getAllCustomersJsonType)
	// http.HandleFunc("/customersXml", getAllCustomersXmlType)

	// // starting servers - Create a server listening on port 8000
	// // Param: 1st: the port details to start server, 2nd: NIl because we are using default mux
	// log.Fatal(http.ListenAndServe("localhost:8000", nil))

	// Above code peice is default Mux, now we define mux

	// Registering routes/endpoints to DefaultServe Mux- All URLs will be handled by this function
	// Param: 1st: pattern for route or endpoint; 2nd: handler func response we want to send to client

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customersJson", getAllCustomersJsonType)
	router.HandleFunc("/customersXml", getAllCustomersXmlType)
	//any string as customer id
	router.HandleFunc("/customers/{customer_id}", getCustomer)
	//integers as customer id
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)

	// starting servers - Create a server listening on port 8000
	// Param: 1st: the port details to start server, 2nd: NIl because we are using default mux
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	//
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["customer_id"])
}
