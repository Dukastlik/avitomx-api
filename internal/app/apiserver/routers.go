package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/add/{link}", s.handleAdd()) // TODO POST request
}

func (s *APIServer) handleAdd() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		go AddFile()
		vars := mux.Vars(r)
		link := vars["link"]
		io.WriteString(w, link)
	}
}
