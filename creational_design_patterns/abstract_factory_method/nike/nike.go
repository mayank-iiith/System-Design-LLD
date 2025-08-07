package nike

import "lld/creational_design_patterns/abstract_factory_method/products"

const (
	nikeLogo = "nike"
)

type Nike struct {
}

// makeShirt implements sports.SportsFactory.
func (n *Nike) MakeShirt(size int) products.Shirt {
	return newNikeShirt(nikeLogo, size)
}

// makeShoe implements sports.SportsFactory.
func (n *Nike) MakeShoe(size int) products.Shoe {
	return newNikeShoe(nikeLogo, size)
}
