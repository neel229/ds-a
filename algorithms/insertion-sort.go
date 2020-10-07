package main

import "fmt"

func (l list) insertionSort() {
	// Compute the length of the list. Will be used as a parameter to stop the for loop.
	size := len(l)
	for i := 1; i < size; i++ {
		key := l[i] // Seperates the sorted array from the unsorted array.
		j := i - 1

		for j >= 0 && l[j] > key {
			l[j+1] = l[j]
			j--
		}
		l[j+1] = key
	}

	fmt.Printf("Sorted list is %v\n", l)
}
