package model

import (
	"errors"
)

// Product model
type Product struct {
	ID         int     `json:"-"`
	MerchantID int     `json:"merchID"`
	OfferID    int     `json:"offerID"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	Available  bool    `json:"available"`
}

func (p *Product) Validate() error {
	if p.Price < 0 || p.Quantity < 0 {
		return errors.New("Negative price or quantity")

	}
	return nil
}
