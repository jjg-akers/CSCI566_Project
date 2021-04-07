package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jjg-akers/csci556_project/registrar"
)

type VoterRegistrationHandler struct {
	registrar.Registrar
}

func (h *VoterRegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println("voter regristrar hit")
	// log.Println(r.URL)
	// enableCors(&w)

	switch r.Method {
	case http.MethodGet:
		log.Println(r.Method)
		log.Println(r.URL.Query().Get("verificationcode"))
	case http.MethodPost:
		log.Println(r.Method)
		log.Println(r.Body)
		// log.Println(r.ParseForm())
		// log.Println(r.URL.Query().Post("fname"))
		// log.Println(r.URL.Query().Post("lname"))
		// log.Println(r.URL.Query().Post("idnumber"))
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		log.Println(r.Method)
		log.Println(r.URL.Query())
		// log.Println("error message")
		// Give an error message.
	}

	// log.Println(r.Method)
	// log.Println(r.URL)

	fmt.Fprint(w, "ok")
}

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }
