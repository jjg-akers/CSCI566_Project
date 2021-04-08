package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/jjg-akers/csci556_project/registrar"
)

type VoterRegistrationHandler struct {
	registrar.Registrar
}

type Voter struct{
	First_name string `schema:"first_name"`
	Last_name string `schema:"last_name"`
	Idnumber string `schema:"idnumber"`
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
		// log.Println(r.Method)
		// log.Println(r.Body)
		d := schema.NewDecoder()
		log.Println("processing post")

		err := r.ParseForm()
		if err != nil{
			log.Println("parse form failed")
			w.WriteHeader(http.StatusBadRequest)
			return
	
		}
	
		voter := &Voter{}

		if err = d.Decode(voter, r.PostForm); err != nil {
			log.Println("err decoding post form: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		v, err := json.MarshalIndent(voter, "", "  ")
		if err != nil{
			log.Println("err marshalling")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println("voter: ", string(v))

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
