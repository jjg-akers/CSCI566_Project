package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jjg-akers/csci556_project/registrar"
)

type VoterRegistrationHandler struct{
	registrar.Registrar
}

func (h *VoterRegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println("voter regristrar hit")

	fmt.Fprint(w, "ok")
}