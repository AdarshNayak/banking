package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// last json tag for tagging in json -> useful for encoding and decoding
type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	// To send response to client
	// Param: 1st: iowriter, 2nd: data to be sent to io writer
	fmt.Fprintf(w, "Hello World!!")
}

// handler func to be called on endpoint customers -> to give list of cuostmers in json fomrat
func getAllCustomersJsonType(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"RamyaRao", "Bangalore", "560048"},
		{"Ashish", "Bangalore", "560048"},
		{"Juyal", "New Delhi", "560048"},
	}

	/* Encode the customer list in json fomrat*/
	// This to show the customer list in proper json fomrat
	w.Header().Add("Content-Type", "application/json")
	//Encoding
	json.NewEncoder(w).Encode(customers)
}

// handler func to be called on endpoint customers -> to give list of cuostmers in json fomrat
func getAllCustomersXmlType(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"RamyaRao", "Bangalore", "560048"},
		{"Ashish", "Bangalore", "560048"},
		{"Juyal", "New Delhi", "560048"},
	}

	// To support content type request from postman requesting for particular encode type
	if r.Header.Get("Content-Type") == "application/xml" {
		/* Encode the customer list in XML format*/
		// This to show the customer list in proper XML fomrat
		w.Header().Add("Content-Type", "application/xml")
		//Encoding
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
