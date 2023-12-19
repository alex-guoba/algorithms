// https://www.hackerrank.com/challenges/three-month-preparation-kit-find-the-running-median/problem
//  - 第一感觉是使用快速排序后直接结算中位数即可。但是testcase中输入数组超过1w时会计算超时

package main

import "fmt"

func runningMedian(a []int32) []float64 {
	// Write your code here
	sorted := make([]int32, len(a))
	result := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		// cur := qsort(a[0:i+1])
		// pos := bsearch(sorted, a[i])

		// copy tail, may be overlaped
		j := 0
		for j = i - 1; j >= 0; j-- {
			if sorted[j] < a[i] {
				j++ // withdraw to the correct
				break
			}
			// move node
			sorted[j+1] = sorted[j]
		}
		// header
		if j < 0 {
			j = 0
		}
		sorted[j] = a[i]

		if (i+1)%2 == 1 {
			result[i] = float64(sorted[(i+1)/2])
		} else {
			result[i] = float64(sorted[i/2]+sorted[(i+1)/2]) / 2
		}
		fmt.Println(result[i], sorted)
	}
	return result
}

func main() {
	a := []int32{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	v := runningMedian(a)
	fmt.Println(v)
}
