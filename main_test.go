package main

import (
	"reflect"
	"testing"
)

func TestGroupIsAesthetic(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  bool
	}{
		"bell curve":          {[]int{1, 4, 1}, true},
		"inverted bell curve": {[]int{3, 1, 7}, true},
		"positive slope":      {[]int{1, 2, 3}, false},
		"negative slope":      {[]int{3, 2, 1}, false},
		"flat":                {[]int{4, 4, 4}, false},
	}

	for name, test := range tests {
		got := groupIsAesthetic(test.input)
		if got != test.want {
			t.Errorf("%s: expected: %v, got: %v", name, test.want, got)
		}
	}
}

func TestDeleteAndAppend(t *testing.T) {
	tests := map[string]struct {
		slice       []int
		deleteIndex int
		appendValue int
		want        []int
	}{
		"delete first":  {[]int{1, 4, 1}, 0, 3, []int{4, 1, 3}},
		"delete second": {[]int{1, 4, 1}, 1, 3, []int{1, 1, 3}},
		"delete third":  {[]int{1, 4, 1}, 2, 3, []int{1, 4, 3}},
	}

	for name, test := range tests {
		got := deleteAndAppend(test.slice, test.deleteIndex, test.appendValue)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s: expected: %v, got: %v", name, test.want, got)
		}
	}
}

func TestPossibleWays(t *testing.T) {
	tests := map[string]struct {
		trees    []int
		nextTree int
		want     int
	}{
		"three ways": {[]int{3, 4, 1}, 3, 3},
		"two ways":   {[]int{1, 4, 4}, 3, 2},
		"no way":     {[]int{1, 2, 3}, 3, 0},
	}

	for name, test := range tests {
		got, _ := possibleWays(test.trees, test.nextTree)
		if got != test.want {
			t.Errorf("%s: expected: %v, got: %v", name, test.want, got)
		}
	}
}

func TestGroupTrees(t *testing.T) {
	tests := map[string]struct {
		trees          []int
		firstTreeIndex int
		want           []int
	}{
		"first tree":  {[]int{3, 4, 1, 5}, 0, []int{3, 4, 1}},
		"second tree": {[]int{3, 4, 1, 5}, 1, []int{4, 1, 5}},
		"third tree":  {[]int{3, 4, 1, 5}, 2, []int{}},
		"fourth tree": {[]int{3, 4, 1, 5}, 3, []int{}},
	}

	for name, test := range tests {
		got := groupTrees(test.trees, test.firstTreeIndex)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s: expected: %v, got: %v", name, test.want, got)
		}
	}
}

func TestSolution(t *testing.T) {
	tests := map[string]struct {
		trees []int
		want  int
	}{
		"three ways":           {[]int{3, 4, 5, 3, 7}, 3},
		"two ways":             {[]int{3, 4, 4, 3, 7}, 2},
		"already aesthetic":    {[]int{1, 3, 1, 2}, 0},
		"can not be aesthetic": {[]int{1, 2, 3, 4}, -1},
	}

	for name, test := range tests {
		got := Solution(test.trees)
		if got != test.want {
			t.Errorf("%s: expected: %v, got: %v", name, test.want, got)
		}
	}
}
