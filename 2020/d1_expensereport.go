package main

import (
	"fmt"
)

func FindMatchingPair(numbers []int64, result int64) (int64, error) {
	lookup := make(map[int64]bool)
	for _, n := range numbers {
		if _, ok := lookup[result-n]; ok {
			return n * (result - n), nil
		}
		lookup[n] = true
	}
	return int64(0), fmt.Errorf("Not found")
}

func FindMatchingTriplet(numbers []int64, result int64) (int64, error) {
	lookup := make(map[int64]bool)
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if _, ok := lookup[result-numbers[i]-numbers[j]]; ok {
				return (numbers[i] * numbers[j] * (result - numbers[i] - numbers[j])), nil
			}
			lookup[numbers[i]] = true
			lookup[numbers[j]] = true
		}
	}
	return int64(0), fmt.Errorf("Not found")
}
