package main

import "fmt"

func InsertSort(s []int) {
	for i := 1; i < len(s); i++ {
		fmt.Println(s)
		for j := i; j > 0 && s[j-1] > s[j]; j-- {
			fmt.Printf("Move %v to the right by 1 ", s[j-1])
			fmt.Println()
			s[j], s[j-1] = s[j-1], s[j]
			fmt.Println(s)
			fmt.Println()
		}
	}
}

func main() {
	s := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	InsertSort(s)
}
