package nike

import (
	"fmt"
	"lld/creational_design_patterns/abstract_factory_method/products"
)

type nikeShoe struct {
	logo string
	size int
}

func newNikeShoe(logo string, size int) products.Shoe {
	return &nikeShoe{
		logo: logo,
		size: size,
	}
}

func (n *nikeShoe) GetLogoAndType() string {
	return fmt.Sprintf("NikeShoe: Logo: %s, Size: %d", n.logo, n.size)
}
