package board

//go:generate mockgen -destination=../../mocks/mock_board.go -package=mocks github.com/thomaslitton/boggle-my-mind/pkg/board BoardWalker

import (
    "github.com/thomaslitton/boggle-my-mind/pkg/word"
)

// Note:  Board should be immutable once created.
type Board struct {
  Rows int
  Columns int
  Letters []rune
}

type BoardWalker interface {
    Walk(*Board, chan word.Word)
}

type BoardWalkerImpl struct {
    paths []word.Path
}

func NewBoardWalker(rows int, columns int) (BoardWalker) {
    paths := CalculatePaths(rows, columns)

    return BoardWalkerImpl{paths: paths}
}

func (b BoardWalkerImpl) Walk(board *Board, foundWords chan word.Word) {
    for _, path := range b.paths {
        w := []rune{}
        for _, index := range path {
            w = append(w, board.Letters[index])
        }

        foundWords <- word.Word{Path: path, Value: string(w)}
    }

    close(foundWords)
}
