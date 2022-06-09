package game_test

import(
    "testing"
    "os"
    "fmt"
    "strings"
    "log"

    "github.com/tomlitton/boggle-my-mind/pkg/game"
    "github.com/tomlitton/boggle-my-mind/pkg/word"
    "github.com/tomlitton/boggle-my-mind/pkg/board"
)

func TestFullGame(t *testing.T) {
    log.Printf("Starting test")
    boardLetters := "naigorlydaisoiuf"
    gameBoard := board.Board{
        Rows: 4,
        Columns: 4,
        Letters: []rune(boardLetters),
    }

    validWords, parseErr := parseValidWords(boardLetters)
    log.Printf("Parsed valid words")

    if parseErr != nil {
        t.Fatalf("Parsing valid words failed: %v", parseErr)
    }

    log.Printf("Getting board walker")
    boardWalker := board.NewBoardWalker(gameBoard.Rows, gameBoard.Columns)
    verifier, dictErr := word.NewWordVerifier("../../assets/dictionary.txt")

    if dictErr != nil {
        t.Fatalf("Failed to load dictionary %v", dictErr)
    }

    log.Printf("Processing game board")
    words, err := game.ProcessGameBoard(&gameBoard, boardWalker, verifier)
    log.Printf("Done")

    if err != nil {
        t.Fatalf("Failed to process game: %v", err)
    }

    for _, w := range validWords {
        if !containsWords(words, w) {
            t.Fatalf("word not found %s", w)
        }
    }
}

// Note:  Very similar functions exist in board_test and can be consolidated if and when necessary.
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
