package apiserver

import (
	"encoding/json"
	"net/http"
)

// AddRequest ...
type AddRequest struct {
	MerchantId int
	FileLink   string
}

func parseRequest(r *http.Request) (*AddRequest, error) {
	var a *AddRequest
	err := json.NewDecoder(r.Body).Decode(&a)
	//fmt.Println(a.FileLink, a.MerchantId)
	if err != nil {
		return nil, err
	}
	return a, nil
}
