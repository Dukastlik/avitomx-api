package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/Dukastlik/avitomx-api.git/internal/app/model"
)

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/add", s.handleAdd()).Methods("POST")
	s.router.HandleFunc("/stat", s.handleStat()).Methods("GET")
}

func (s *APIServer) handleAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		addreq, err := parseAddRequest(r)
		if err != nil {
			http.Error(w, "Invalid data proceeded", http.StatusBadRequest)
			return
		}

		respStruct, err := AddFile(addreq, s)
		if err != nil {
			http.Error(w, "Invalid file proceeded", http.StatusBadRequest)
			return
		}

		js, err := respStruct.ToJson()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func (s *APIServer) handleStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		statreq, err := model.ParseStatRequest(r)
		if err != nil {
			http.Error(w, "Invalid data proceeded", http.StatusBadRequest)
			return
		}
		products, err := GetStat(statreq, s)
		if err != nil {
			http.Error(w, "Invalid data proceeded", http.StatusBadRequest)
			return
		}
		js, err := json.Marshal(products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
