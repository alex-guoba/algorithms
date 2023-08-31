package main

import "fmt"

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // 选择数组的最后一个元素作为主元
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // 交换元素
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // 将主元放到正确的位置
	return i + 1
}

func main() {
	data := []int{9, 5, 2, 7, 12, 6, 3, 1, 8, 4}
	fmt.Println("Before sorting:", data)
	quickSort(data, 0, len(data)-1)
	fmt.Println("After sorting:", data)
}
