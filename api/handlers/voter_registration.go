package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jjg-akers/csci556_project/registrar"
	"github.com/jjg-akers/csci556_project/util"
)

type VoterRegistrationHandler struct {
	registrar.Registrar
}

type Voter struct{
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Idnumber string `json:"idnumber"`
}

func (h *VoterRegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if origin := r.Header.Get("Origin"); origin != "" {
        w.Header().Set("Access-Control-Allow-Origin", origin)
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }

	    // Stop here if its Preflighted OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}

	log.Println("voter regristrar hit")


	switch r.Method {
	case http.MethodGet:
		log.Println(r.Method)
		log.Println(r.URL.Query().Get("verificationcode"))
	case http.MethodPost:

		// ----- Uncomment to read raw body
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
        //     log.Printf("Error reading body: %v", err)
        //     http.Error(w, "can't read body", http.StatusBadRequest)
        //     return
        // }
		//		fmt.Println("body: ", string(body))
		// --------------

		
		log.Println("processing post")
	
		voter := &Voter{}

		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&voter)
		if err != nil {
			log.Println("err decoding")
			w.WriteHeader(http.StatusInternalServerError)		}
	
		v, err := json.MarshalIndent(voter, "", "  ")
		if err != nil{
			log.Println("err marshalling")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Println("voter: ", string(v))

		// hash values
		str := voter.First_name + voter.Last_name + voter.Idnumber

		hash := util.StringToKeccak256(str)

		str = util.KeccakToString(hash[:])

		fmt.Println("voter id string: ", str)

		w.Header().Set("Content-Type", "application/json")
		
		fmt.Fprintf(w, `{"voter_id":%q}`, str)

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

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, "ok")
}

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }
