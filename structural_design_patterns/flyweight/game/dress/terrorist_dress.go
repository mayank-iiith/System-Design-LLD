package dress

// terroristDress - Concrete flyweight object
type terroristDress struct {
	color string
}

func newTerroristDress() Dress {
	return &terroristDress{
		color: "red",
	}
}

func (d *terroristDress) GetColor() string {
	return d.color
}
