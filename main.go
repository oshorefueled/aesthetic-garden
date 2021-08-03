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

// deleteAndAppend deletes an element with the index indexToDelete from a slice treeGroup and
// appends a value toAppend to the end of the slice. It then returns the new slice.
func deleteAndAppend(treeGroup []int, indexToDelete, toAppend int) []int {
	t := make([]int, 0)
	t = append(t, treeGroup[:indexToDelete]...)
	t = append(t, treeGroup[indexToDelete+1:]...)
	t = append(t, toAppend)
	return t
}

// possibleWays returns the number of ways one tree can be cut from a group of 3 trees to make
// them aesthetically pleasing. It checks for aesthetic possibility by adding the first tree
// from the next group of trees after removing one tree.
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

	return
}

// groupTrees returns a group of 3 trees from a slice trees starting from an index
// indexOfFirstTreeInGroup
func groupTrees(trees []int, firstTreeIndex int) []int {
	if len(trees)-(firstTreeIndex+3) > -1 {
		return trees[firstTreeIndex : firstTreeIndex+3]
	}

	return []int{}
}

func cutDownTree(trees []int, treeIndex int) []int {
	return append(trees[:treeIndex], trees[treeIndex+1:]...)
}

func Solution(A []int) int {
	var ways int

loop:
	for i := 0; i < len(A); i++ {
		if groupedTrees := groupTrees(A, i); len(groupedTrees) == 3 {
			if groupIsAesthetic(groupedTrees) {
				continue
			}

			// ways > 0 means a tree has already been cut down
			// only one tree can be legally cut down to achieve aesthetic desire, return -1
			if ways > 0 {
				return canNotBeAesthetic
			}

			w, toCut := possibleWays(groupedTrees, A[i+3])
			if w == 0 {
				// if there are no ways to resolve aesthetic issue return -1
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
	fmt.Println(Solution([]int{3, 4, 5, 3, 7})) // print 3
}
