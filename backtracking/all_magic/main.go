// 生成n*n魔方数组的暴力解法
// 使用back-tracking，列出所有可能的组合
// https://josephjsf2.github.io/hackerrank/2020/01/23/Forming-a-Magic-Square.html
//
//链接：https://leetcode.cn/problems/permutations/solutions/857631/dai-ma-sui-xiang-lu-dai-ni-xue-tou-hui-s-mfrp/
//

package main

import (
	"fmt"
	"math"
)

var (
	res  [][]int
	path []int
	st   []bool // state的缩写
)

func permute(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(nums))
	st = make([]bool, len(nums))
	dfs(nums, 0)
	return res
}

func is_magic(nums []int) bool {
	depth := int(math.Sqrt(float64(len(nums))))
	if depth*depth != len(nums) {
		return false
	}

	sumHist := 0

	// row
	for i := 0; i < depth; i++ {
		sum := 0
		for j := 0; j < depth; j++ {
			sum += nums[i*depth+j]
		}
		if sumHist == 0 {
			sumHist = sum
		} else if sum != sumHist {
			return false
		}
	}

	// column
	for i := 0; i < depth; i++ {
		sum := 0
		for j := 0; j < depth; j++ {
			sum += nums[j*depth+i]
		}

		if sumHist == 0 {
			sumHist = sum
		} else if sum != sumHist {
			return false
		}
	}

	// square
	sumB2T := 0
	sumT2B := 0
	for i := 0; i < depth; i++ {
		sumT2B += nums[i*depth+i]
		sumB2T += nums[(i+1)*depth-i-1]
	}
	if sumHist != sumB2T || sumHist != sumT2B {
		return false
	}

	return true
}

func dfs(nums []int, cur int) {
	if cur == len(nums) {
		if !is_magic(path) {
			return
		}
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !st[i] {
			path = append(path, nums[i])
			st[i] = true
			dfs(nums, cur+1)
			st[i] = false
			path = path[:len(path)-1]
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//arr := []int{1, 2, 3, 4}
	permute(arr)
	for i, a := range res {
		fmt.Println(i, ":", a)
	}
}
