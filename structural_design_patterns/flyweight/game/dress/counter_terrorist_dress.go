package dress

// counterTerroristDress - Concrete flyweight object
type counterTerroristDress struct {
	color string
}

func newCounterTerroristDress() Dress {
	return &counterTerroristDress{
		color: "red",
	}
}

func (d *counterTerroristDress) GetColor() string {
	return d.color
}
