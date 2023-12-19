package main

import "fmt"

/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */
func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	result := []int32{}
	count := 0
	backtrack(s, d, int(m), result, 0, &count)
	return int32(count)
}

func sumof(a []int32) int32 {
	var t int32
	for _, v := range a {
		t += v
	}
	return t
}

func backtrack(s []int32, sum int32, level int, current []int32, start int, count *int) {

	if len(current) == level {
		if sumof(current) == sum {
			fmt.Println(current)
			*count = *count + 1
		}
		return
	}

	// TODO: 剪枝

	for i := start; i < len(s); i++ {
		current = append(current, s[i])
		backtrack(s, sum, level, current, i+1, count)
		current = current[0 : len(current)-1]
	}
}

func main() {
	s := []int32{2, 2, 1, 3, 2}
	fmt.Println(birthday(s, 4, 2))
}
