package products

import (
	"github.com/Dukastlik/avitomx-api.git/internal/app/model"
)

// ProductRepository ...
type ProductRepository struct {
	products *Products
}

// Create ...
func (r *ProductRepository) Create(p *model.Product) (*model.Product, error) {
	if err := r.products.db.QueryRow(
		"INSERT INTO products (offerID, name, price, quantity, avaliable) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		p.OfferID,
		p.Name,
		p.Price,
		p.Quantity,
		p.Available,
	).Scan(&p.ID); err != nil {
		return nil, err
	}
	return p, nil
}
