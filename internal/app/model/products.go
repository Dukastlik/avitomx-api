package model

// Product model
type Product struct {
	ID        int
	OfferID   int
	Name      string
	Price     int
	Quantity  int
	Available bool
}
