package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n")
		matrixLen, _ := strconv.ParseInt(text, 10, 64)

		switch matrixLen {
		case 0:
			os.Exit(0)
		default:
			matrix := readMatrix(reader, matrixLen)
			// fmt.Printf("matrix is %v\n", matrix)
			handleMatrix(matrix, matrixLen)
		}
	}
}

const (
	splitChars = " "

	okStatus         = "OK"
	repairableStatus = "REPAIRABLE"
	errorStatus      = "ERROR"
)

func readMatrix(reader *bufio.Reader, matrixLen int64) [][]int64 {
	result := make([][]int64, matrixLen)

	for i := int64(0); i < matrixLen; i++ {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n")

		numChars := strings.Split(text, splitChars)
		nums := make([]int64, matrixLen)

		for i, numChar := range numChars {
			nums[i], _ = strconv.ParseInt(numChar, 10, 64)
		}

		result[i] = nums
	}

	return result
}

func handleMatrix(matrix [][]int64, matrixLen int64) {

	sumOfRows := make([]int64, matrixLen)
	// get sum of rows
	for i := int64(0); i < matrixLen; i++ {
		sum := int64(0)

		for j := int64(0); j < matrixLen; j++ {
			sum = sum + matrix[i][j]
		}

		sumOfRows[i] = sum
	}

	sumOfColumns := make([]int64, matrixLen)
	// get sum of columns
	for j := int64(0); j < matrixLen; j++ {
		sum := int64(0)

		for i := int64(0); i < matrixLen; i++ {
			sum = sum + matrix[i][j]
		}

		sumOfColumns[j] = sum
	}

	status, i, j := judge(sumOfRows, sumOfColumns)
	switch status {

	case okStatus:
		fmt.Println(status)
	case errorStatus:
		fmt.Println(status)
	case repairableStatus:
		fmt.Printf("CHANGE (%v,%v)\n", i+1, j+1)
	}
}

func judge(sumOfRows []int64, sumOfColumns []int64) (string, int, int) {
	numOfOdd1 := 0
	firstOddRow := 0
	for i, sum := range sumOfRows {
		if sum%2 == 1 {
			numOfOdd1++
			firstOddRow = i
		}
	}

	numOfOdd2 := 0
	firstOddColumn := 0
	for i, sum := range sumOfColumns {
		if sum%2 == 1 {
			numOfOdd2++
			firstOddColumn = i
		}
	}

	if numOfOdd1 == 0 && numOfOdd2 == 0 {
		return okStatus, firstOddRow, firstOddColumn
	}

	if numOfOdd1 == 1 && numOfOdd2 == 1 {
		return repairableStatus, firstOddRow, firstOddColumn
	}

	return errorStatus, firstOddRow, firstOddColumn
}
