package products

import (
	"database/sql"
	"fmt"
	"time"

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
	var err error
	fmt.Println("connecting to db on:", p.config.DatabaseURL)
	for i := 0; i < 3; i++ {
		db, err := sql.Open("postgres", p.config.DatabaseURL)
		if err != nil {
			fmt.Printf("Unable to Open DB: %s... Retrying\n", err.Error())
			time.Sleep(time.Second * 2)
		} else if err := db.Ping(); err != nil {
			fmt.Printf("Unable to Open DB: %s... Retrying\n", err.Error())
			time.Sleep(time.Second * 2)
		} else {
			err = nil
			p.db = db
			break
		}

	}
	if err != nil {
		return err
	}
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
