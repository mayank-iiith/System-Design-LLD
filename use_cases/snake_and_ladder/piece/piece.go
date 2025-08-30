package piece

type Piece struct {
	position int
}

func NewPiece() *Piece {
	return &Piece{
		position: 0,
	}
}

func (p *Piece) Position() int {
	return p.position
}

func (p *Piece) SetPosition(position int) {
	p.position = position
}
