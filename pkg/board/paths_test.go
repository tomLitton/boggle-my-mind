package board_test

import (
    "testing"
    "github.com/thomaslitton/boggle-my-mind/pkg/board"
    "github.com/thomaslitton/boggle-my-mind/pkg/word"
)

func Test2X2(t *testing.T) {
    paths := board.CalculatePaths(2,2)

    for _, p := range PATHS2X2 {
        if !containsPath(paths, p) {
            t.Fatalf("Valid path %v not found in paths returned", p)
        }
    }
}

func Test4X4(t *testing.T) {
    paths := board.CalculatePaths(4,4)

    for _, p := range PATH4X4_VALID {
        if !containsPath(paths, p) {
            t.Fatalf("Valid path %v not found in paths returned", p)
        }
    }

    for _, p := range PATH4X4_INVALID {
        if containsPath(paths, p) {
            t.Fatalf("Invalid path %v found in paths returned", p)
        }
    }
}

func Test1X1(t * testing.T) {
    t.Skip("TODO: Implement")
}

func containsPath(pathList []word.Path, path word.Path) (bool) {
    for _, pathToCheck := range pathList {
        if subPath(pathToCheck, path) {
            return true
        }
    }

    return false
}

func subPath(path word.Path, pathSubset word.Path) (bool) {
    if len(path) < len(pathSubset) {
        return false
    }
	for i := 0;i < len(pathSubset);i++ {
	    if path[i] != pathSubset[i] {
	        return false
	    }
	}

	return true
}