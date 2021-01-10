package products

import (
	"github.com/Dukastlik/avitomx-api.git/internal/app/model"
)

// ProductRepository ...
type ProductRepository struct {
	products *Products
}

// Create new instance in db
func (r *ProductRepository) Create(p *model.Product) (*model.Product, error) {
	if err := r.products.db.QueryRow(
		`INSERT INTO products 
		(merchant_id, offer_id, name, price, quantity, avaliable) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id`,
		p.MerchantID,
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

// Update instance in db
func (r *ProductRepository) Update(p *model.Product) (*model.Product, error) {
	if err := r.products.db.QueryRow(
		`UPDATE products 
		SET (merchant_id, offer_id, name, price, quantity, avaliable) = ($1, $2, $3, $4, $5, $6) 
		WHERE merchant_id = $1 AND offer_id = $2 
		RETURNING id`,
		p.MerchantID,
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

// Delete instance from db
func (r *ProductRepository) Delete(p *model.Product) (*model.Product, error) {
	if err := r.products.db.QueryRow(
		"DELETE FROM products WHERE merchant_id = $1 AND offer_id = $2 RETURNING id",
		p.MerchantID,
		p.OfferID,
	).Scan(&p.ID); err != nil {
		return nil, err
	}
	return p, nil
}

// Read rows from db
func (r *ProductRepository) Read(req *model.StatRequest) ([]model.Product, error) {
	rows, err := r.products.db.Query(
		`SELECT * 
		FROM products 
		WHERE 
		($1 = 0 or merchant_id = $1) AND 
		($2 = 0 or offer_id = $2) AND 
		($3 = '' or name SIMILAR TO CONCAT($3, '%'))`,
		req.MerchID,
		req.OfferID,
		req.ProdName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := []model.Product{}
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(
			&p.ID,
			&p.MerchantID,
			&p.OfferID,
			&p.Name,
			&p.Price,
			&p.Quantity,
			&p.Available,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
