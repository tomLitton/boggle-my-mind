package word

//go:generate mockgen -destination=../../mocks/mock_word.go -package=mocks github.com/tomlitton/boggle-my-mind/pkg/word WordVerifier

import (
    "os"
    "strings"
    "errors"
    "fmt"

    "golang.org/x/exp/slices"
)

type Path []int

type WordVerifier interface {
    // Technically this implementation can't fail, but one possible future implementation is to call an API
    // to verify the word.  In that case, an error might be returned.
    Verify(Word) (bool, error)
}

type Word struct {
  Path Path `json:"path"`
  Value string `json:"word"`
}

type WordVerifierImpl struct {
    dictionary []string
}

func NewWordVerifier(dictionaryFile string) (WordVerifier, error) {
    dict, err := parseDictionary(dictionaryFile)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("Failed to parse dictionary: %v", err))
    }

    slices.Sort(dict)
    return WordVerifierImpl{dictionary: dict}, nil
}

func (v WordVerifierImpl) Verify(word Word) (bool, error) {
    _, found := slices.BinarySearch(v.dictionary, word.Value)

    return found, nil
}

func parseDictionary(dictionaryFile string) ([]string, error) {
    contents, err := os.ReadFile(dictionaryFile)
    if err != nil {
        return nil, err
    }

    words := strings.Split(string(contents), "\n")

    for i, w := range words {
        words[i] = strings.TrimSpace(w)
    }

    return words, nil
}
