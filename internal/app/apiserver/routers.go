package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/add", s.handleAdd()) // TODO POST request
}

func (s *APIServer) handleAdd() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		adreq, err := parseRequest(r)
		if err != nil {
			//fmt.Println(err)
			return // TODO ??
		}
		go AddFile(adreq)
		vars := mux.Vars(r)
		link := vars["link"]
		io.WriteString(w, link)
	}
}
