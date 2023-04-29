package Homework

import (
	"sort"
)

func MinSwaps2(A []int) int {

	swaps := 0
	i := 0

	for i < len(A) {
		if A[i] != i {
			A[i], A[A[i]] = A[A[i]], A[i]
			swaps++
		} else {
			i++
		}

	}

	return swaps
}

func MinSwaps2Old(A []int) int {

	positionMap := make(map[int]int)

	// store value and index as key value pair
	for i := range A {
		positionMap[A[i]] = i
	}

	// sort the mao based on key
	keys := make([]int, 0)

	for k, _ := range positionMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	swapCount := 0
	for i := 0; i < len(A); i++ {

		index := positionMap[keys[i]]

		if i != index {
			keys[i], keys[index] = keys[index], keys[i]
			swapCount++
			i--
		}

	}

	return swapCount
}
