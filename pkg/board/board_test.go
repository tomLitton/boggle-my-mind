package board_test

import (
    "testing"
    "os"
    "fmt"
    "strings"

    "github.com/tomlitton/boggle-my-mind/pkg/board"
    "github.com/tomlitton/boggle-my-mind/pkg/word"
)

func TestBoards(t *testing.T) {
    // @TODO: Test additional solved boards
    boardLetters := "naigorlydaisoiuf"
    gameBoard := board.Board{Rows: 4, Columns: 4, Letters: []rune(boardLetters)}
    boardWalker := board.NewBoardWalker(gameBoard.Rows, gameBoard.Columns)

    validWords, err := parseValidWords(boardLetters)
    if err != nil {
        t.Fatalf("Reading valid words failed: %v", err)
    }

    foundWords := make(chan word.Word)
    go boardWalker.Walk(&gameBoard, foundWords)

    words := []word.Word{}
    for w := range foundWords {
        words = append(words, w)
    }

    // ensure the possible words at least contains all the valid words for this board.
    for _, w := range validWords {
        if !containsWords(words, w) {
            t.Fatalf("word not found %s", w)
        }
    }
}

// @TODO: Test boards of other sizes, specifically non-square boards

func parseValidWords(letters string) ([]string, error) {
    contents, err := os.ReadFile(fmt.Sprintf("../../test/fixtures/%s.txt", letters))
    if err != nil {
        return nil, err
    }

    words := strings.Split(string(contents), "\n")

    for i, w := range words {
        words[i] = strings.TrimSpace(w)
    }

    return words, nil
}

func containsWords(words []word.Word, validWord string) (bool) {
    for _, w := range words {
        if w.Value == validWord {
            return true
        }
    }

    return false
}
