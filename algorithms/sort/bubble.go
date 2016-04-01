package main

import "fmt"

func BubbleSort(s []int) {
	fmt.Println("Start BubbleSort:")
	len := len(s) - 1
	swapCounter := 0
	swapped := true
	for swapped {
		swapped = false
		fmt.Print(s)
		fmt.Println(" swapCounter is", swapCounter)
		for i := 0; i < len; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				swapped = true
				swapCounter++
			}
		}
		len--
	}
}

func BubbleSort2(a []int) {
	fmt.Println("Start BubbleSort2:")
	swapCounter := 0
	for i := 0; i <= len(a)-2; i++ {
		fmt.Print(a)
		fmt.Println(" swapCounter is", swapCounter)
		for j := len(a) - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j] //数据交换
				swapCounter++
			}
		}
	}
}

func main() {
	s := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	BubbleSort(s)

	fmt.Println()

	a := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	BubbleSort2(a)
}
