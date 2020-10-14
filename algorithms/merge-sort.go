package main

import "fmt"

func (l list) mergeSort(s, e int) {
	if len(l) < 2 {
		fmt.Println(l)
	}

	mid := len(l) / 2

	if s < e {
		l.mergeSort(s, mid)
		l.mergeSort(mid+1, e)
		
	}
}
