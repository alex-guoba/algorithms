//2.6 - 081. 允许重复选择元素的组合
//给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。
//candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。
//对于给定的输入，保证和为 target 的唯一组合数少于 150 个。
//
// 输入: candidates = [2,3,6,7], target = 7
// 输出: [[7],[2,2,3]]
//
// 输入: candidates = [2,3,5], target = 8
// 输出: [[2,2,2,2],[2,3,3],[3,5]]

package main

import "fmt"

type Array []int

func sum(current Array) int {
	s := 0
	for _, x := range current {
		s += x
	}
	return s
}

func calCandidate(target int, candidates Array, current Array, index int) {
	if sum(current) == target {
		fmt.Println(current)
	}
	//fmt.Println("discard: ", current, " sum: ", sum(current))
	if sum(current) >= target {
		return
	}

	for i := index; i < len(candidates); i++ {
		current = append(current, candidates[i])
		calCandidate(target, candidates, current, i)
		current = current[:len(current)-1]
	}
}

func main() {
	var current Array
	//candidates := []int{2, 3, 6, 7}
	//calCandidate(7, candidates, current, 0)
	candidates := []int{2, 3, 5}
	calCandidate(8, candidates, current, 0)
}
