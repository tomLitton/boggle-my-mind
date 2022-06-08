package board

import (
    "golang.org/x/exp/slices"
)

func CalculateAdjacentIndexes(rows int, columns int) ([][]int) {
    adjacentIndex := make([][]int, rows*columns)
    for r := 0;r < rows;r++ {
        for c :=0;c < columns;c++ {
            index := calculateIndex(r, c, columns)
            adjacentIndex[index] = adjacentIndexesOfIndex(index, r, c, rows, columns)
        }
    }
    return adjacentIndex
}

func adjacentIndexesOfIndex(index, row, column, totalRows, totalColumns int) ([]int) {
    nextRow := row + 1
    prevRow := row - 1

    nextColumn := column + 1
    prevColumn := column - 1

    adjacentIndexes := []int{}

    if nextRow < totalRows {
        adjacentIndexes = append(adjacentIndexes, calculateIndex(nextRow, column, totalColumns))
        if nextColumn < totalColumns {
            adjacentIndexes = append(adjacentIndexes, calculateIndex(nextRow, nextColumn, totalColumns))
        }

        if prevColumn >= 0 {
            adjacentIndexes = append(adjacentIndexes, calculateIndex(nextRow, prevColumn, totalColumns))
        }
    }

    if prevRow >= 0 {
        adjacentIndexes = append(adjacentIndexes, calculateIndex(prevRow, column, totalColumns))
        if nextColumn < totalColumns {
            adjacentIndexes = append(adjacentIndexes, calculateIndex(prevRow, nextColumn, totalColumns))
        }

        if prevColumn >= 0 {
            adjacentIndexes = append(adjacentIndexes, calculateIndex(prevRow, prevColumn, totalColumns))
        }
    }

    if nextColumn < totalColumns {
        adjacentIndexes = append(adjacentIndexes, calculateIndex(row, nextColumn, totalColumns))
    }

    if prevColumn >= 0 {
        adjacentIndexes = append(adjacentIndexes, calculateIndex(row, prevColumn, totalColumns))
    }

    // Note:  This isn't necessary, but makes debugging and testing easier.
    slices.Sort(adjacentIndexes)

    return adjacentIndexes
}

func calculateIndex(row, column, totalColumns int) (int) {
    return row * totalColumns + column
}