package piece

type PieceType string

const (
	PieceTypeO     PieceType = "O"
	PieceTypeX     PieceType = "X"
	PieceTypeEmpty PieceType = "-"
)

func (p PieceType) String() string {
	return string(p)
}
