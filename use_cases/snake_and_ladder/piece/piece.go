package piece

type Piece struct {
	currentPosition int
}

func NewPiece() *Piece {
	return &Piece{
		currentPosition: 0,
	}
}

func (p *Piece) GetCurrentPosition() int {
	return p.currentPosition
}

func (p *Piece) SetNewPosition(position int) {
	p.currentPosition = position
}
