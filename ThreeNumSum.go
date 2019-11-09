package main

import (
	"fmt"
	"sort"
)

// ThreeNumSum is an algo-- given an input array
// and a target, determine if three numbers sum
// to the target
func ThreeNumSum(inputArr []int, target int) bool {

	// init three pointers, one that moves across
	// the inputArry, then two at each index
	// by ratcheting inward

	// have to sort the array for the ratcheting
	// to work

	sort.Ints(inputArr)

	for pos, num := range inputArr {

		left := pos + 1
		right := len(inputArr) - 1

		for left < right {

			sum := num + inputArr[left] + inputArr[right]

			fmt.Println(sum, num, left, right)

			if sum == target {

				return true
			}

			if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	return false
}

// func main() {

// 	inputArr := []int{11, 22, 110, 120, 33, 44}
// 	target := 99

// 	fmt.Println(ThreeNumSum(inputArr, target))
// }
