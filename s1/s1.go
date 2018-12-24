package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n")

		switch text {
		case "0":
			os.Exit(0)
		default:
			p := calculateP(text)
			fmt.Printf("%v\n", p)
		}
	}
}

func calculateP(input string) int {
	inputLen := len(input)
	// fmt.Printf("inputLen = %v\n", inputLen)

	for sliceLen := 1; sliceLen <= (inputLen / 2); sliceLen++ {

		if inputLen%sliceLen != 0 {
			continue
		}

		firstSlice := input[0:sliceLen]

		// fmt.Printf("sliceLen = %v, firstSlice is %v\n", sliceLen, firstSlice)

		if checkIsRepeatable(input, firstSlice) {
			// fmt.Printf("found, input = %v, firstSlice is %v\n", input, firstSlice)
			return inputLen / sliceLen
		}
	}

	return 1
}

func checkIsRepeatable(input, firstSlice string) bool {

	len1 := len(input)
	len2 := len(firstSlice)

	for s := 0; s < len1; s = s + len2 {
		if s+len2 > len1 {
			return false
		}

		curSlice := input[s:(s + len2)]
		if curSlice != firstSlice {
			return false
		}
	}

	return true
}
