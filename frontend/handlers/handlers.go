package handlers

import (
	"net/http"
)


type IndexHandler struct {

}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "/static/index.html")

	http.FileServer(http.Dir("templates")).ServeHTTP(w, r)

		// http.Handle("/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

}