package sports

import "lld/creational_design_patterns/abstract_factory_method/products"

type SportsFactory interface {
	MakeShirt(size int) products.Shirt
	MakeShoe(size int) products.Shoe
}
