package adidas

import (
	"fmt"
	"lld/creational_design_patterns/abstract_factory_method/products"
)

type adidasShirt struct {
	logo string
	size int
}

func newAdidasShirt(logo string, size int) products.Shirt {
	return &adidasShirt{
		logo: logo,
		size: size,
	}
}

func (n *adidasShirt) GetLogoAndType() string {
	return fmt.Sprintf("AdidasShirt: Logo: %s, Size: %d", n.logo, n.size)
}
