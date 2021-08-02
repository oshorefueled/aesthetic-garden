package main

import (
	"fmt"
)

const (
	canNotBeAesthetic = -1
)

// groupIsAesthetic checks if a group of 3 trees is aesthetically pleasing. It returns true if the group is
// aesthetically pleasing, false if it is not. It also returns false if the length of its argument has a length
// less than 3
//
// a group of 3 trees is aesthetically pleasing when the second tree is either the
// highest or lowest of the three
func groupIsAesthetic(group []int) bool {
	if len(group) < 3 {
		// a group of 3 trees is needed to confirm aesthetic desire
		return false
	}
	middleTreeIsLowest := group[1] < group[0] && group[1] < group[2]
	middleTreeIsHighest := group[1] > group[0] && group[1] > group[2]
	if middleTreeIsLowest || middleTreeIsHighest {
		return true
	}

	return false
}

func deleteAndAppend(treeGroup []int, indexToDelete, toAppend int) []int {
	t := make([]int, 0)
	t = append(t, treeGroup[:indexToDelete]...)
	t = append(t, treeGroup[indexToDelete+1:]...)
	t = append(t, toAppend)
	return t
}

func possibleWays(group []int, treeAfter int) (ways int, treeToCut int) {
	checkAndUpdateWays := func(treeIndex int) {
		if groupIsAesthetic(deleteAndAppend(group, treeIndex, treeAfter)) {
			ways++
			treeToCut = treeIndex
		}
	}

	for k := range group {
		checkAndUpdateWays(k)
	}

	// if there are no ways to resolve aesthetic issue return -1
	if ways == 0 {
		ways = canNotBeAesthetic
	}

	return
}

func treeCanBeGrouped(trees []int, indexOfFirstTreeInGroup int) ([]int, bool) {
	if len(trees)-(indexOfFirstTreeInGroup+3) > -1 {
		return trees[indexOfFirstTreeInGroup : indexOfFirstTreeInGroup+3], true
	}

	return []int{}, false
}

func cutDownTree(trees []int, treeIndex int) []int {
	return append(trees[:treeIndex], trees[treeIndex+1:]...)
}

func Solution(A []int) int {
	var ways int

loop:
	for i := 0; i < len(A); i++ {
		if groupedTrees, canGroup := treeCanBeGrouped(A, i); canGroup {
			if groupIsAesthetic(groupedTrees) {
				continue
			}

			// ways > 0 means a tree has already been cut down
			// only one tree can be cut down to achieve aesthetic desire, return -1
			if ways > 0 {
				return canNotBeAesthetic
			}

			w, toCut := possibleWays(groupedTrees, A[i+3])
			if w == -1 {
				return canNotBeAesthetic
			}

			ways += w
			A = cutDownTree(A, i+toCut)
			goto loop
		}
	}

	return ways
}

func main() {
	fmt.Println(Solution([]int{1, 4, 4, 1, 9, 9}))
}
