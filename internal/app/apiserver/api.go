package apiserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Dukastlik/avitomx-api.git/internal/app/model"
)

func printTime() {
	dt := time.Now()
	fmt.Println("time: ", dt.String())
}

func AddFile(a *AddRequest, s *APIServer) (*AddResponse, error) {

	// getting xlsx file from url
	resp, err := http.Get(a.FileLink)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// copying from response to byte slice
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// parsing xlsx to slice of products
	products, invalid, err := ParseXlsx(data)
	if err != nil {
		return nil, err
	}

	deleted, updated, created := 0, 0, 0
	//checking if structure is valid
	for _, p := range products {
		p.MerchantID = a.MerchantId

		if p.Validate() != nil {
			invalid++
		} else { // if structure is valid adding it to db, or updating exisiting row
			if p.Available != true {
				_, err := s.products.Product().Delete(p) // Deleting item from db
				if err != nil {
					invalid++
				} else {
					deleted++
				}
			} else {
				_, err := s.products.Product().Update(p) // Updating db item
				if err != nil {
					_, err := s.products.Product().Create(p) // Creating new item
					if err != nil {
						invalid++
					} else {
						created++
					}
				} else {
					updated++
				}
			}
		}
	}
	// creating AddResponse struct
	addresp := &AddResponse{
		Created: created,
		Updated: updated,
		Deleted: deleted,
		Invalid: invalid,
	}

	return addresp, nil
}

func GetStat(st *model.StatRequest, s *APIServer) ([]model.Product, error) {
	fmt.Println("merch id : ", st.MerchID)
	fmt.Println("offer id : ", st.OfferID)
	fmt.Println("name: ", st.ProdName)

	products, err := s.products.Product().Read(st)
	if err != nil {
		fmt.Println(err)
	}

	return products, nil
}
