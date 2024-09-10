package algo

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}

func BubbleSortDemo() {
	a := []int{82, 33, 91, 9, 1, 553, 31, 3923}
	sorted := bubbleSort(a)
	for v := range sorted {
		fmt.Println(sorted[v])
	}
}
