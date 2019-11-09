package main

import (
	"math"
	"sort"
	"testing"
)

/*

Write a function that takes in two non-empty arrays of integers.
The function should find the pair of numbers
(one from the first array, one from the second array)
whose absolute difference is closest to zero.
The function should return an array containing these two numbers,
with the number from the first array in the first position.
Assume that there will only be one pair of numbers with
the smallest difference.

*/

// SmallestDifference is an algo
func SmallestDifference(left []int, right []int) [2]int {

	// sort each array, then initialize a pointer at
	// the beginning of each- compare the ints at each
	// pointer, and increment the index of the array
	// with the smaller int

	sort.Ints(left)
	sort.Ints(right)

	leftpt := 0
	rightpt := 0

	returnArr := [2]int{left[0], right[0]}

	for (leftpt < (len(left))) && (rightpt < (len(right))) {

		currentClosest := math.Abs(float64(left[0] - right[0]))

		thisLeft := left[leftpt]
		thisRight := right[rightpt]

		currentDiff := math.Abs(float64(thisLeft - thisRight))

		if currentDiff < currentClosest {

			returnArr[0] = thisLeft
			returnArr[1] = thisRight
		}

		if thisLeft == thisRight {

			return returnArr
		} else if thisLeft < thisRight {

			leftpt++
		} else {

			rightpt++
		}

	}

	return returnArr
}

type testInputs struct {
	left   []int
	right  []int
	expect [2]int
}

var tests = []testInputs{

	{[]int{0}, []int{0}, [2]int{0, 0}},

	{[]int{1}, []int{1}, [2]int{1, 1}},

	{[]int{5, 1, 2, 3, 4}, []int{5, 6, 7, 8}, [2]int{5, 5}},

	{[]int{2, 1, 2, 3, 4}, []int{5, 6, 7, 8}, [2]int{4, 5}},

	{[]int{2, 1, 2, 3, 4}, []int{53, 64, 75, 86}, [2]int{4, 53}},
}

// TestSmallestDifference is testing my algo
func TestSmallestDifference(t *testing.T) {

	// fmt.Println("TestSmallestDifference")

	for _, inputs := range tests {

		result := SmallestDifference(inputs.left, inputs.right)

		if result != inputs.expect {

			t.Error("Expected", inputs.expect, "got", result)
		}

	}
}
