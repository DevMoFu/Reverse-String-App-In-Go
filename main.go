package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func splitChrs(str string) []string {
	if len(str) == 0 || len(str) == 1 {
		fmt.Println("String of 1 chr or nil provided\nEnter a valid string and try again")
		os.Exit(1)
	}
	var splitChrs []string
	splitChrs = strings.Split(str, "")

	var reversedChr []string
	var lastIndex int = len(str) - 1

	for i, j := 0, lastIndex; j >= 0; i, j = i+1, j-1 {
		reversedChr = append(reversedChr, splitChrs[j])
	}

	return reversedChr
}

func main() {
	userSting := flag.String("String", "Defualt Value", "String to mod")
	flag.Parse()

	var reverseTextSlice = splitChrs(*userSting)
	var reverseText string = strings.Join(reverseTextSlice, "")

	fmt.Printf("\nReversed String: '%v' \n\n", reverseText)
}
