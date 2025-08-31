package piece

import "errors"

func GetPiece(pt string) (PieceType, error) {
	switch pt {
	case "X":
		return PieceTypeX, nil
	case "O":
		return PieceTypeO, nil
	default:
		return PieceTypeEmpty, errors.New("invalid pieceType")
	}
}
