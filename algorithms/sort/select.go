package main

import "fmt"

func SelectSort(s []int) {
	fmt.Println("Start SelectSort:")
	len := len(s)
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if s[j] < s[i] {
				s[i], s[j] = s[j], s[i]
			}
		}
		fmt.Println(s)
	}
}

func main() {
	s := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	fmt.Println("Source:", s)
	fmt.Println()
	SelectSort(s)
}
