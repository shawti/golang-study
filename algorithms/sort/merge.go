package main

import "fmt"

func MergeSort(s []int, lo, hi int) {
	aux := make([]int, len(s))
	MergeSortSub(s, aux, lo, hi)
}

func MergeSortSub(s, aux []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := (lo + hi) >> 1
	MergeSortSub(s, aux, lo, mid)
	MergeSortSub(s, aux, mid+1, hi)
	Merge(s, aux, lo, mid, hi)
}

func Merge(s, aux []int, lo, mid, hi int) {
	for k := lo; k <= hi; k++ {
		aux[k] = s[k]
	}
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			s[k] = aux[j]
			j++
		} else if j > hi {
			s[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			s[k] = aux[j]
			j++
		} else {
			s[k] = aux[i]
			i++
		}
	}
}

func main() {
	s := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	MergeSort(s, 0, len(s)-1)
	fmt.Println(s)
}
