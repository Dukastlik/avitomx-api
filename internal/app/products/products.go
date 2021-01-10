package products

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// Products ...
type Products struct {
	config   *Config
	db       *sql.DB
	prodRepo *ProductRepository
}

// New ...
func New(config *Config) *Products {
	return &Products{
		config: config,
	}
}

// Open ...
func (p *Products) Open() error {

	db, err := sql.Open("postgres", p.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	p.db = db

	return nil
}

// Close ...
func (p *Products) Close() {
	p.db.Close()
}

// Product ...
func (p *Products) Product() *ProductRepository {
	if p.prodRepo != nil {
		return p.prodRepo
	}

	p.prodRepo = &ProductRepository{
		products: p,
	}
	return p.prodRepo
}

//products.Product.Create
