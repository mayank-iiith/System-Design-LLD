package adidas

import (
	"fmt"
	"lld/creational_design_patterns/abstract_factory_method/products"
)

type adidasShoe struct {
	logo string
	size int
}

func newAdidasShoe(logo string, size int) products.Shoe {
	return &adidasShirt{
		logo: logo,
		size: size,
	}
}

func (n *adidasShoe) GetLogoAndType() string {
	return fmt.Sprintf("AdidasShoe: Logo: %s, Size: %d", n.logo, n.size)
}
