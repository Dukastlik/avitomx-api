package model

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
)

type StatRequest struct {
	MerchID  int    `schema:"merchID"`
	OfferID  int    `schema:"offerID"`
	ProdName string `schema:"prod_name"`
}

func ParseStatRequest(r *http.Request) (*StatRequest, error) {
	//var a *StatRequest
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	/*err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		return nil, err
	}*/
	a := new(StatRequest)
	if err := schema.NewDecoder().Decode(a, r.Form); err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("NEW PARSED NAME", a.ProdName)
	return a, nil
}
