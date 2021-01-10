package apiserver

import (
	"errors"

	"github.com/Dukastlik/avitomx-api.git/internal/app/model"
	"github.com/tealeg/xlsx/v3"
)

func ModelGeter(r *xlsx.Row) (*model.Product, error) {
	ofID, err1 := r.GetCell(0).Int()
	name := r.GetCell(1).String()
	price, err2 := r.GetCell(2).Float()
	quantity, err3 := r.GetCell(3).Int()
	avaliable := r.GetCell(4).Bool()

	if err1 != nil || err2 != nil || err3 != nil {
		return nil, errors.New("Invalid row")
	}

	productmodel := &model.Product{
		OfferID:   ofID,
		Name:      name,
		Price:     price,
		Quantity:  quantity,
		Available: avaliable,
	}
	return productmodel, nil
}

func ParseXlsx(data []byte) ([]*model.Product, int, error) {

	wb, err := xlsx.OpenBinary(data)

	if err != nil {
		panic(err) //TODO
	}

	var parsedProd []*model.Product
	var invalidCount = 0
	//for every sheet in doc
	for _, sh := range wb.Sheets {
		// for every row on a sheet
		for i := 0; i < sh.MaxRow; i++ {

			r, _ := sh.Row(i)
			if err != nil {
				panic(err) //TODO
			}
			//creating Product structure from a row data
			prod, err := ModelGeter(r)
			if err != nil {
				invalidCount++
			} else {
				parsedProd = append(parsedProd, prod)
			}

		}
	}
	return parsedProd, invalidCount, nil
}
