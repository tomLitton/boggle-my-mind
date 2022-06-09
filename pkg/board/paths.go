package board

import (
    "golang.org/x/exp/slices"

    "github.com/tomlitton/boggle-my-mind/pkg/word"
)

var adjacentIndexes [][]int

func CalculatePaths(rows int, columns int) ([]word.Path) {
    adjacentIndexes = CalculateAdjacentIndexes(rows, columns)

    count := rows * columns
    allPaths := []word.Path{}

    // @TODO:  add rule for min 3 letter words
    for i := 0;i < count;i++ {
        thisPath := word.Path{i}
        allPaths = append(allPaths, thisPath)
        usedIndexes := make([]bool, count)
        usedIndexes[i] = true
        for _, nextIndex := range adjacentIndexes[i] {
            // Note:  Doing a fork and join here would be a relatively easy optimization
            subPath := calculatePaths(thisPath, nextIndex, usedIndexes)
            allPaths = append(allPaths, subPath...)
        }
    }

    return allPaths
}

func calculatePaths(startingPath word.Path, index int, usedIndexes []bool) ([]word.Path) {
    if usedIndexes[index] {
        return []word.Path{}
    }

    allPaths := []word.Path{}

    newPath := slices.Clone(startingPath)
    usedIndexes = slices.Clone(usedIndexes)

    newPath = append(newPath, index)
    allPaths = append(allPaths, newPath)
    usedIndexes[index] = true

    for _, nextIndex := range adjacentIndexes[index] {
        subPath := calculatePaths(newPath, nextIndex, usedIndexes)
        allPaths = append(allPaths, subPath...)
    }

    return allPaths
}