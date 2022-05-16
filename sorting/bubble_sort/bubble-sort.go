package bubble_sort

import "fmt"

func BubbleSort(arr [8]int) [8]int {
	fmt.Println("==== before sort =====")
	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {

				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
		}
	}
	fmt.Println("===== after sort ==== ")

	return arr
}
