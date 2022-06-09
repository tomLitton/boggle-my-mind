package game_test

import (
    "testing"
    "errors"

    "github.com/golang/mock/gomock"

    "github.com/tomlitton/boggle-my-mind/mocks"
    "github.com/tomlitton/boggle-my-mind/pkg/game"
    "github.com/tomlitton/boggle-my-mind/pkg/word"
    "github.com/tomlitton/boggle-my-mind/pkg/board"
)

var WORD1 = word.Word {
    Path: word.Path{1, 2, 3},
    Value: "Word",
}

var WORD2 = word.Word{
    Path: word.Path{5, 6, 7},
    Value: "Word",
}

var NOT_WORD1 = word.Word{
    Path: word.Path{8, 9, 10},
    Value: "NotAWord",
}

var NOT_WORD2 = word.Word{
    Path: word.Path{11, 12, 13},
    Value: "NotAWord",
}

var BOARD = board.Board{
    Rows: 2,
    Columns: 2,
    Letters: []rune{'a','b','c','d'},
}

func TestWhenSomeAreWords(t *testing.T) {
    runTest(t, []word.Word{WORD1, NOT_WORD1, WORD2})
}

func TestWhenAllAreWords(t *testing.T) {
    runTest(t, []word.Word{WORD1, WORD2})
}

func TestWhenNoPotentialWords(t *testing.T) {
    runTest(t, []word.Word{})
}

func TestWhenNoValidWords(t *testing.T) {
    runTest(t, []word.Word{NOT_WORD1, NOT_WORD2})
}

// Note that this needs to be corrected.  Verify failures _should_ return an error
func TestWhenVerifyFails(t *testing.T) {
    possibleWords := []word.Word{WORD1, NOT_WORD1, WORD2}
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockBoardWalker := mocks.NewMockBoardWalker(mockCtrl)
    mockWordVerifier := mocks.NewMockWordVerifier(mockCtrl)

    mockBoardWalker.EXPECT().Walk(&BOARD, gomock.Any()).Do(mockWalk(possibleWords))

    for _, w := range possibleWords {
        if w.Value == "NotAWord" {
            mockWordVerifier.EXPECT().Verify(w).Return(false, errors.New("Testing"))
        } else {
            mockWordVerifier.EXPECT().Verify(w).Return(w.Value == "Word", nil)
        }
    }

    _, err := game.ProcessGameBoard(&BOARD, mockBoardWalker, mockWordVerifier)

    if err != nil {
        t.Fatalf("Game returned an error %v", err)
    }
}

func runTest(t *testing.T, possibleWords []word.Word) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockBoardWalker := mocks.NewMockBoardWalker(mockCtrl)
    mockWordVerifier := mocks.NewMockWordVerifier(mockCtrl)

    mockBoardWalker.EXPECT().Walk(&BOARD, gomock.Any()).Do(mockWalk(possibleWords))

    for _, w := range possibleWords {
        mockWordVerifier.EXPECT().Verify(w).Return(w.Value == "Word", nil)
    }

    results, err := game.ProcessGameBoard(&BOARD, mockBoardWalker, mockWordVerifier)

    if err != nil {
        t.Fatalf("Game returned an error %v", err)
    }

    for _, w := range possibleWords {
        if w.Value == "Word" &&  !containsWord(w, results) {
            t.Fatalf("Expected result not found.  %v not in list %v", w, results)
        } else if w.Value != "Word" &&  containsWord(w, results) {
            t.Fatalf("Invalid result found.  %v in list %v", w, results)
        }
    }
}

func mockWalk(words []word.Word) (func(*board.Board, chan word.Word)) {
    return func(board *board.Board, foundWords chan word.Word) {
        for _, w := range words {
            foundWords <- w
        }

        close(foundWords)
    }
}

func containsWord(word word.Word, wordList []word.Word) (bool) {
    for _, w := range wordList {
        if w.Value == word.Value {
            return true
        }
    }
    return false
}