package api

import (
    "log"
    "errors"
    "net/http"
    "fmt"
    "encoding/json"

    "github.com/gorilla/mux"

    "github.com/tomlitton/boggle-my-mind/pkg/game"
    "github.com/tomlitton/boggle-my-mind/pkg/board"
    "github.com/tomlitton/boggle-my-mind/pkg/word"
)

const ROWS = 4
const COLUMNS = 4

var boardWalker board.BoardWalker
var verifier word.WordVerifier

type Answer struct {
    Letters string `json:"letters"`
    Words []word.Word `json:"words"`
}

func LoadData(dictionaryFile string) {
    var dictErr error
    boardWalker = board.NewBoardWalker(ROWS, COLUMNS)
    verifier, dictErr = word.NewWordVerifier(dictionaryFile)

    if dictErr != nil {
        log.Fatalf("Failed to load dictionary")
    }
}

func HandleBoard(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    gameBoard, boardErr := createBoard(vars)
    if boardErr != nil {
        writeBadRequest(w, boardErr.Error())
        return
    }

    log.Printf("Processing board")
    words, err := game.ProcessGameBoard(gameBoard, boardWalker, verifier)
    log.Printf("Done processing board")

    if err != nil {
        writeInternalError(w, fmt.Sprintf("Error processing game: %v", err))
        return
    }

    writeAnswer(w, Answer{Letters: vars["letters"], Words: words})
}

func createBoard(vars map[string]string) (*board.Board, error) {
    rowsStr := vars["rows"]
    columnsStr := vars["columns"]

    if rowsStr != "4" || columnsStr != "4" {
        return nil, errors.New("Only 4 X 4 boards are supported")
    }

    letters := vars["letters"]

    if len(letters) != (ROWS * COLUMNS) {
        return nil, errors.New("Wrong number of letters")
    }

    return &board.Board{Rows: ROWS, Columns: COLUMNS, Letters: []rune(letters)}, nil
}

func writeAnswer(w http.ResponseWriter, answer Answer) {
    body, jsonErr := json.Marshal(answer)
    if jsonErr != nil {
        writeInternalError(w, fmt.Sprintf("Error writing answer: %v", jsonErr))
        return
    }

    _, writeErr := w.Write([]byte(body))

    if writeErr != nil {
        log.Printf("Error writing answer: %v", writeErr)
    }
}

func writeBadRequest(w http.ResponseWriter, msg string) {
    w.WriteHeader(http.StatusBadRequest)
    writeErrorResponse(w, msg)
}

func writeInternalError(w http.ResponseWriter, msg string) {
    log.Print(msg)
    w.WriteHeader(http.StatusInternalServerError)
    writeErrorResponse(w, msg)
}

func writeErrorResponse(w http.ResponseWriter, msg string) {
    _, writeErr := w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", msg)))
    if writeErr != nil {
        log.Printf("Failed to write response: %v", writeErr)
    }
}

func Health(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("OK"))
}