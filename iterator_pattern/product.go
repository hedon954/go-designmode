package main

type Product struct {
	Name  string
	Price float64
}

type ProductList struct {
	Products []*Product
}

// CreateIterator returns an iterator for ProductList
func (pl *ProductList) CreateIterator() Iterator {
	return &ProductIterator{
		Index:    0,
		Products: pl.Products,
	}
}

// ProductIterator is the iterator for product list
type ProductIterator struct {
	Index    int
	Products []*Product
}

func (p *ProductIterator) HasNext() bool {
	return p.Index < len(p.Products)
}

func (p *ProductIterator) Next() interface{} {
	defer func() {
		p.Index++
	}()
	return p.Products[p.Index]
}
