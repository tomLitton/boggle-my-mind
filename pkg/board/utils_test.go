package board_test

import (
    "testing"
    "golang.org/x/exp/slices"
     "github.com/tomlitton/boggle-my-mind/pkg/board"
)

func TestAdjacentCellsWith2X2(t *testing.T) {
    adjacentIndexes := board.CalculateAdjacentIndexes(2,2)

    for i :=0;i < 4;i++ {
        expectedList := slices.Delete([]int{0,1,2,3}, i, i+1)
        if !slices.Equal(adjacentIndexes[i], expectedList) {
            t.Fatalf("Adjacent index is not correct for index %d: %v 1= %v", i, adjacentIndexes[i], expectedList)
        }
    }
}

func TestAdjacentCellsWith1X1(t *testing.T) {
    t.Skip("TODO: Implement test")
}
