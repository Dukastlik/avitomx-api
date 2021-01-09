package apiserver

import (
	"github.com/Dukastlik/avitomx-api.git/internal/app/model"
	"github.com/tealeg/xlsx/v3"
)

//Xlsx parser
/*func cellVisitor(c *xlsx.Cell) error {

	colnum, _ := c.GetCoordinates()

	value, err := c.FormattedValue()
	if err != nil {
		return err
	}
		switch colnum {
		case 0:
			p.OfferID = value
		case 1:
			p.Name = value
		case 2:
			p.Price = value
		case 3:
			p.Quantity = value
		case 4:
			p.Available = value
		}

	return err
} */

/*func rowVisitor(r *xlsx.Row) model.Product {

	productmodel := model.Product{
		OfferID:   r.GetCell(0),
		Name:      r.GetCell(1),
		Price:     r.GetCell(2),
		Quantity:  r.GetCell(3),
		Available: r.GetCell(4),
	}

	return productmodel
}*/

func ModelGeter(r *xlsx.Row) model.Product {
	ofID, _ := r.GetCell(0).Int()
	name := r.GetCell(1).String()
	price, _ := r.GetCell(2).Int()
	quantity, _ := r.GetCell(3).Int()
	avaliable := r.GetCell(4).Bool()

	productmodel := model.Product{
		OfferID:   ofID,
		Name:      name,
		Price:     price,
		Quantity:  quantity,
		Available: avaliable,
	}
	//fmt.Println(productmodel.Name, productmodel.Price, productmodel.Quantity)
	return productmodel
}

func ParseXlsx(data []byte) {

	wb, err := xlsx.OpenBinary(data)

	if err != nil {
		panic(err)
	}

	for _, sh := range wb.Sheets {

		for i := 0; i < sh.MaxRow; i++ {

			r, _ := sh.Row(i)
			if err != nil {
				panic(err) //TODO
			}
			ModelGeter(r)

		}
	}
}
