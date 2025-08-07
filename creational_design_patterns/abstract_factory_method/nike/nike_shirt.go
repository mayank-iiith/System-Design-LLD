package nike

import (
	"fmt"
	"lld/creational_design_patterns/abstract_factory_method/products"
)

type nikeShirt struct {
	logo string
	size int
}

func newNikeShirt(logo string, size int) products.Shirt {
	return &nikeShirt{
		logo: logo,
		size: size,
	}
}

func (n *nikeShirt) GetLogoAndType() string {
	return fmt.Sprintf("NikeShirt: Logo: %s, Size: %d", n.logo, n.size)
}
