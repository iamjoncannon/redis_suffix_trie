package main

import (
	"math"
)

// TwoNumSum is an algorithm
// given an input array and a target, determine if
// two numbers in the array sum to target
func TwoNumSum(inputArr []int, target int) bool {

	cache := make(map[int]bool)

	for _, num := range inputArr {

		diff := int(math.Abs(float64(target - num)))

		if cache[diff] == true {

			return true
		}

		cache[num] = true
	}

	return false
}

// func main() {

// 	sampleArr := []int{11, 22, 33, 44}
// 	target := 55

// 	fmt.Println(TwoNumSum(sampleArr, target))
// }
