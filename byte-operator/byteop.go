package main

import "fmt"

// Roy 想要提高他编程比赛的打字速度，所以，他的朋友建议他反复输入这个句子：“The quick brown fox jumps over the lazy dog”，因为这个句子是全字母短句，（全字母短句中每个字母都出现至少一次）。
//在输入这个句子很多次之后，Roy 觉得很无聊，所以他开始寻找其它的全字母短句。
//给定一个句子 , 告诉 Roy 这个句子是不是一个全字母短句。
//输入格式
//输入只有一行字符串s
//数据约束
// - s长度最多10000
//它可能包括空格，小写字母及大写字母。小写字母跟大写字母可以当成同一个字母。

//输出格式
//如果是全字母短句， 输出 pangram
//否则输出 not pangram.

func pangrams(s string) string {
	// Write your code here
	var result int32 = 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		var step byte
		if c >= 'a' && c <= 'z' {
			step = byte('z') - c + 1
		} else if c >= 'A' && c <= 'Z' {
			step = byte('Z') - c + 1
		}

		if step >= 1 && step <= 26 {
			result = result | (0x1 << int32(step-1))
		}
	}

	dest := int32((0xffffff << 2) | 0xf)

	if result == dest {
		return "pangram"
	}
	return "not pangram"
}
func main() {

	fmt.Println(pangrams("We promptly judged antique ivory buckles for the next prize"))
	// calc reversal
	//var d uint8 = 2
	//var xd uint32 = 1024
	//var xxd int64 = 4294967294
	//fmt.Printf("%08b\n", d)
	//fmt.Printf("%08b\n", ^d)
	//fmt.Printf("%32b\n", xd)
	//fmt.Printf("%32b\n", ^xd)
	//fmt.Printf("%64b\n", xxd)
	//fmt.Printf("%64b\n", ^uint64(xxd))
	//fmt.Printf("%64b\n", int64(^uint64(xxd)))
	//fmt.Printf("%64b\n", int64(^uint64(xxd)&0xFFFFFFFF))

	//fmt.Printf("%32b\n", 0x1<<16)
}
