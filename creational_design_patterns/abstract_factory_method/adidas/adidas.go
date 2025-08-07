package adidas

import (
	"lld/creational_design_patterns/abstract_factory_method/products"
)

const (
	adidasLogo = "adidas"
)

type Adidas struct {
}

func (a *Adidas) MakeShirt(size int) products.Shirt {
	return newAdidasShirt(adidasLogo, size)
}

func (a *Adidas) MakeShoe(size int) products.Shoe {
	return newAdidasShoe(adidasLogo, size)
}
