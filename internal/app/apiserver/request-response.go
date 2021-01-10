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

func parseAddRequest(r *http.Request) (*AddRequest, error) {
	var a *AddRequest
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

type AddResponse struct {
	Created int `json:"Products created"`
	Updated int `json:"Products updated"`
	Deleted int `json:"Products deleted"`
	Invalid int `json:"Products invalid"`
}

func (ar *AddResponse) ToJson() ([]byte, error) {
	return json.Marshal(*ar)
}
