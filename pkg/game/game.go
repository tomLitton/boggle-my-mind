package game

import (
    "log"
    "sync"

    "github.com/tomlitton/boggle-my-mind/pkg/word"
    "github.com/tomlitton/boggle-my-mind/pkg/board"
)

const WORD_VERIFIER_THREADS = 20

func ProcessGameBoard(board *board.Board, walker board.BoardWalker, verifier word.WordVerifier)  ([]word.Word, error) {
    verifiedWord := make(chan word.Word)
    foundWords := make(chan word.Word)

    // We can fork this as well if necessary
    go walker.Walk(board, foundWords)

    var wg sync.WaitGroup
    for i := 0;i < WORD_VERIFIER_THREADS;i++ {
        wg.Add(1)
        go verifyWords(foundWords, verifiedWord, &wg, verifier)
    }

    go func() {
        wg.Wait()
        close(verifiedWord)
    }()

    words := []word.Word{}
    for r := range verifiedWord {
        words = append(words, r)
    }

    return words, nil
}

func verifyWords(foundWords chan word.Word, verifiedWord chan word.Word, wg *sync.WaitGroup, verifier word.WordVerifier) {
    defer wg.Done()

    for w := range foundWords {
        isWord, err := verifier.Verify(w)
        if err != nil {
            log.Printf("Failed to verify word %s: %q", w.Value, err)
        }

        if isWord {
            verifiedWord <- w
        }
    }
}
