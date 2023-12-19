package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	s1, _ := reader.ReadString('\n')
	s1 = strings.Trim(s1, "\n")

	s2, _ := reader.ReadString('\n')
	s2 = strings.Trim(s2, "\n")

	result := ""
	for {
		if len(s1) == 0 && len(s2) == 0 {
			break
		}
		var c1 byte = '0'
		if len(s1) > 0 {
			c1 = s1[len(s1)-1]
			s1 = s1[0 : len(s1)-1]
		}
		var c2 byte = '0'
		if len(s2) > 0 {
			c2 = s2[len(s2)-1]
			s2 = s2[0 : len(s2)-1]
		}

		if c1 == c2 {
			result = "0" + result
		} else {
			result = "1" + result
		}
	}
	fmt.Printf("%s", result)
}
