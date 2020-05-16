package main

import (
	"flag"
	"fmt"
	"strings"
)

// SplitChrs to split characters
func SplitChrs(str string) []string {
	/* if len(str) == 0 || len(str) == 1 {
		fmt.Println("String of 1 chr or nil provided\nEnter a valid string and try again")
		os.Exit(1)
	} */
	splitChrs := strings.Split(str, "")
	fmt.Println(splitChrs)
	/* var reversedChr []string

	var lastIndex int = len(str) - 1

	for i, j := 0, lastIndex; i > lastIndex || j < 0; i, j = i+1, j-1 {
		reversedChr[i] = splitChrs[j]
	}

	return reversedChr */
	return []string{"s", "t"}
}

func main() {
	var str = flag.String("String", "", "String to modify")
	var reverseTextSlice = SplitChrs(*str)
	var reverseText string = strings.Join(reverseTextSlice, "")

	fmt.Printf("Reversed string: %v", reverseText)
}
